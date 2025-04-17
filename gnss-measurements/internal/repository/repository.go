package statistics_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	measurements_domain "gnss-radar/gnss-measurements/internal"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type RadarRequest struct {
	RadarX         float64  `json:"radar_x"`
	RadarY         float64  `json:"radar_y"`
	RadarZ         float64  `json:"radar_z"`
	InspectionTime uint64   `json:"inspection_time"`
	TLEFile        string   `json:"tle_file"`
	SatellitesName []string `json:"satellites_name"`
}

type PgxIFace interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type MeasurementsRepo struct {
	pool   PgxIFace
	minio  *minio.Client
	logger *logrus.Logger
}

func NewMeasurementsRepo(pool PgxIFace, m *minio.Client, logger *logrus.Logger) *MeasurementsRepo {
	return &MeasurementsRepo{pool: pool, minio: m, logger: logger}
}

func (mr *MeasurementsRepo) GetEphemeris(ctx context.Context, req measurements_domain.PaginatedRequest) ([]measurements_domain.EphemerisFileMeta, uint64, error) {

	var total uint64

	countQuery := `SELECT COUNT(*) FROM file_meta;`
	err := mr.pool.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get total count")
	}

	fileMetaQuery := `
		SELECT
			filename,
			created_at
		FROM file_meta 
		ORDER BY created_at DESC
		LIMIT $1
		OFFSET $2;
	`
	var ephemeris []measurements_domain.EphemerisFileMeta

	offset := (req.Page - 1) * req.Size

	rows, err := mr.pool.Query(ctx, fileMetaQuery, req.Size, offset)
	if err != nil {
		return ephemeris, total, errors.Wrap(err, "failed to get ephemeris info")
	}

	defer rows.Close()

	for rows.Next() {
		var timeStamp time.Time
		var fileMeta measurements_domain.EphemerisFileMeta
		err := rows.Scan(
			&fileMeta.Name,
			&timeStamp,
		)
		if err != nil {
			return nil, total, errors.Wrap(err, "failed to scan row")
		}
		fileMeta.Datetime = timeStamp.Format("2006-01-02 15:04:05")
		ephemeris = append(ephemeris, fileMeta)
	}

	if err := rows.Err(); err != nil {
		return nil, total, errors.Wrap(err, "error during rows iteration")
	}

	return ephemeris, total, nil
}

func (mr *MeasurementsRepo) UploadEphemeris(ctx context.Context, req measurements_domain.EphemerisToLoad) error {
	bucketName := os.Getenv("EPHEMERIS_BUCKET_NAME")
	if _, err := mr.minio.PutObject(
		ctx,
		bucketName,
		req.MinioName,
		req.Payload,
		req.PayloadSize,
		minio.PutObjectOptions{ContentType: req.ContentType},
	); err != nil {
		return err
	}

	writeFileMeta := `
	INSERT INTO file_meta(filename, minio_name) values 
	($1, $2);
	`

	_, err := mr.pool.Exec(ctx, writeFileMeta, &req.Name, &req.MinioName)
	if err != nil {
		err := mr.minio.RemoveObject(ctx, bucketName, req.MinioName, minio.RemoveObjectOptions{ForceDelete: true})
		if err != nil {
			return errors.Wrap(err, "ALARM! File added, but it is not in database")
		}

		return err
	}

	return nil
}

func (mr *MeasurementsRepo) GetSatellitesCoordinates(ctx context.Context) (measurements_domain.Satellites, error) {
	requestBody := RadarRequest{
		RadarX:         56.4475,
		RadarY:         37.423056,
		RadarZ:         0.5,
		SatellitesName: []string{},
	}

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return measurements_domain.Satellites{}, errors.Wrap(err, "failed to load Moscow location")
	}

	now := time.Now().In(loc)

	formattedInt, err := strconv.Atoi(now.Format("20060102150405"))
	if err != nil {
		return measurements_domain.Satellites{}, errors.Wrap(err, "failed to format time")
	}

	requestBody.InspectionTime = uint64(formattedInt)

	var tleName string

	query := `SELECT minio_name FROM file_meta ORDER BY created_at DESC;`
	err = mr.pool.QueryRow(ctx, query).Scan(&tleName)
	if err != nil {
		return measurements_domain.Satellites{}, errors.Wrap(err, "failed to get minio name")
	}

	requestBody.TLEFile = tleName

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return measurements_domain.Satellites{}, err
	}

	//fmt.Println(requestBody)

	satellites_addr := os.Getenv("SATELLITES_ADDR")

	url := fmt.Sprintf("%s/api/v1/satellites/now", satellites_addr)

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return measurements_domain.Satellites{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return measurements_domain.Satellites{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return measurements_domain.Satellites{}, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var result measurements_domain.Satellites
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return measurements_domain.Satellites{}, err
	}

	return result, nil
}
