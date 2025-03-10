package statistics_repository

import (
	"context"
	measurements_domain "gnss-radar/gnss-measurements/internal"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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
		ORDER BY created_at
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

func (mr MeasurementsRepo) UploadEphemeris(ctx context.Context, req measurements_domain.EphemerisToLoad) error {
	bucketName := os.Getenv("EPHEMERIS_BUCKET_NAME")
	if _, err := mr.minio.PutObject(
		ctx,
		bucketName,
		req.Name,
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

	err := mr.pool.QueryRow(ctx, writeFileMeta, &req.Name, &req.MinioName)
	if err != nil {
		err := mr.minio.RemoveObject(ctx, bucketName, req.Name, minio.RemoveObjectOptions{ForceDelete: true})
		if err != nil {
			return errors.Wrap(err, "ALARM! File added, but it is not in database")
		}

		return err
	}

	return nil
}
