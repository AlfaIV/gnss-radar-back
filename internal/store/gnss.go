package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type IGnssStore interface {
	ListGnssCoords(ctx context.Context, filter ListGnssCoordsFilter) ([]*model.GnssCoords, error)
	CreateDevice(ctx context.Context, params CreateDeviceParams) (*model.Device, error)
	UpdateDevice(ctx context.Context, params UpdateDeviceParams) (*model.Device, error)
	ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error)
	DeleteDevice(ctx context.Context, filter DeleteDeviceFilter) error
	CreateTask(ctx context.Context, params CreateTaskParams) (*model.Task, error)
	UpdateTask(ctx context.Context, params UpdateTaskParams) (*model.Task, error)
	DeleteTask(ctx context.Context, filter DeleteTaskFilter) error
	ListTask(ctx context.Context, filter ListTasksFilter) ([]*model.Task, error)
	ListSatellites(ctx context.Context, filter ListSatellitesFilter) ([]*model.SatelliteInfo, error)
	CreateSatellite(ctx context.Context, params CreateSatelliteParams) (*model.SatelliteInfo, error)
	RinexList(ctx context.Context) ([]*model.RinexResults, error)
	AddSpectrum(ctx context.Context, spectrumReq model.SpectrumRequest) error
	AddPower(ctx context.Context, powerReq model.PowerRequest) error
	AddPairMeasurement(ctx context.Context, pairMeasurementReq model.PairMeasurementRequest) error
	ListMeasurements(ctx context.Context, measurementReq model.MeasurementsFilter) ([]*model.Measurement, error)
	CompareDeviceToken(ctx context.Context, deviceTokenReq string) error
}

type GnssStore struct {
	storage *Storage
}

type ListDeviceFilter struct {
	Ids       []string `db:"id"`
	Names     []string `db:"name"`
	Tokens    []string `db:"token"`
	Paginator model.Paginator
}

func (g *GnssStore) ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error) {
	query := g.storage.Builder().
		Select("id, name, token, description, x, y, z, created_at").
		From(deviceTable)

	if len(filter.Names) > 0 {
		query = query.Where(sq.Eq{"name": filter.Names})
	}
	if len(filter.Tokens) > 0 {
		query = query.Where(sq.Eq{"token": filter.Tokens})
	}
	if len(filter.Ids) > 0 {
		query = query.Where(sq.Eq{"id": filter.Ids})
	}
	if filter.Paginator.Page != 0 {
		query = query.Offset(filter.Paginator.Page)
	}
	if filter.Paginator.PerPage != 0 {
		query = query.Limit(filter.Paginator.PerPage)
	}

	var devices []*model.Device
	if err := g.storage.db.Selectx(ctx, &devices, query); err != nil {
		return nil, postgresError(err)
	}

	return devices, nil
}

func NewGnssStore(storage *Storage) *GnssStore {
	return &GnssStore{
		storage: storage,
	}
}

type ListGnssCoordsFilter struct {
	X         float64
	Y         float64
	Z         float64
	Paginator model.Paginator
}

func (g *GnssStore) ListGnssCoords(ctx context.Context, filter ListGnssCoordsFilter) ([]*model.GnssCoords, error) {
	query := g.storage.Builder().
		Select("id, satellite_id, x, y, z, created_at").
		From(gnssTable)

	if filter.X != 0 {
		query = query.Where(sq.Eq{"x": filter.X})
	}
	if filter.Y != 0 {
		query = query.Where(sq.Eq{"y": filter.Y})
	}
	if filter.Z != 0 {
		query = query.Where(sq.Eq{"z": filter.Z})
	}
	if filter.Paginator.Page != 0 {
		query = query.Offset(filter.Paginator.Page)
	}
	if filter.Paginator.PerPage != 0 {
		query = query.Limit(filter.Paginator.PerPage)
	}

	var coords []*model.GnssCoords
	if err := g.storage.db.Selectx(ctx, &coords, query); err != nil {
		return nil, postgresError(err)
	}

	return coords, nil
}

type CreateDeviceParams struct {
	Name        string  `db:"name"`
	Token       string  `db:"token"`
	Description *string `db:"description"`
	Coords      model.Coords
}

func (g *GnssStore) CreateDevice(ctx context.Context, params CreateDeviceParams) (*model.Device, error) {
	query := g.storage.Builder().
		Insert(deviceTable).
		SetMap(map[string]any{
			"name":        params.Name,
			"token":       params.Token,
			"description": params.Description,
			"x":           params.Coords.X,
			"y":           params.Coords.Y,
			"z":           params.Coords.Z,
		}).
		Suffix("RETURNING id, name, token, description, x, y, z, created_at")

	var device model.Device
	if err := g.storage.db.Getx(ctx, &device, query); err != nil {
		return nil, postgresError(err)
	}

	return &device, nil
}

type UpdateDeviceParams struct {
	Id          string  `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`
	Coords      model.Coords
}

func (g *GnssStore) UpdateDevice(ctx context.Context, params UpdateDeviceParams) (*model.Device, error) {
	query := g.storage.Builder().
		Update(deviceTable).
		Where(sq.Eq{"id": params.Id}).
		SetMap(map[string]any{
			"name":        params.Name,
			"description": params.Description,
			"x":           params.Coords.X,
			"y":           params.Coords.Y,
			"z":           params.Coords.Z,
		}).
		Suffix("RETURNING id, name, token, description, x, y, z, created_at")

	var device model.Device
	if err := g.storage.db.Getx(ctx, &device, query); err != nil {
		return nil, postgresError(err)
	}

	return &device, nil
}

type CreateTaskParams struct {
	SatelliteId  string
	DeviceId     string
	Title        string
	Description  *string
	SignalType   model.SignalType
	GroupingType model.GroupingType
	StartAt      time.Time
	EndAt        time.Time
}

func (g *GnssStore) CreateTask(ctx context.Context, params CreateTaskParams) (*model.Task, error) {
	query := g.storage.Builder().
		Insert(taskTable).
		SetMap(map[string]any{
			"satellite_id":  params.SatelliteId,
			"device_id":     params.DeviceId,
			"title":         params.Title,
			"description":   params.Description,
			"signal_type":   params.SignalType,
			"grouping_type": params.GroupingType,
			"start_at":      params.StartAt,
			"end_at":        params.EndAt,
		}).
		Suffix("RETURNING id, title, description, satellite_id, device_id, signal_type, grouping_type, start_at, end_at, created_at")

	var task model.Task
	if err := g.storage.db.Getx(ctx, &task, query); err != nil {
		return nil, postgresError(err)
	}

	return &task, nil
}

type DeleteTaskFilter struct {
	Id string
}

func (g *GnssStore) DeleteTask(ctx context.Context, filter DeleteTaskFilter) error {
	query := g.storage.Builder().
		Delete(taskTable).
		Where(sq.Eq{"id": filter.Id})

	if filter.Id != "" {
		query = query.Where(sq.Eq{"id": filter.Id})
	}

	if _, err := g.storage.db.Execx(ctx, query); err != nil {
		return postgresError(err)
	}

	return nil
}

type UpdateTaskParams struct {
	Id           string
	Title        string
	Description  *string
	SatelliteID  string
	SignalType   model.SignalType
	GroupingType model.GroupingType
	StartAt      time.Time
	EndAt        time.Time
}

func (g *GnssStore) UpdateTask(ctx context.Context, params UpdateTaskParams) (*model.Task, error) {
	query := g.storage.Builder().
		Update(taskTable).
		Where(sq.Eq{"id": params.Id}).
		SetMap(map[string]any{
			"satellite_id":  params.SatelliteID,
			"title":         params.Title,
			"description":   params.Description,
			"signal_type":   params.SignalType,
			"grouping_type": params.GroupingType,
			"start_at":      params.StartAt,
			"end_at":        params.EndAt,
		}).
		Suffix("RETURNING id, title, description, satellite_id, signal_type, grouping_type, start_at, end_at, created_at")

	var task model.Task
	if err := g.storage.db.Getx(ctx, &task, query); err != nil {
		return nil, postgresError(err)
	}

	return &task, nil
}

type ListTasksFilter struct {
	Ids          []string
	SatelliteIds []string
	DeviceId     []string
	SignalType   []model.SignalType
	GroupingType []model.GroupingType
	StartAt      *time.Time
	EndAt        *time.Time
	Paginator    model.Paginator
}

func (g *GnssStore) ListTask(ctx context.Context, filter ListTasksFilter) ([]*model.Task, error) {
	query := g.storage.Builder().
		Select("id, title, description, satellite_id, device_id, signal_type, grouping_type, start_at, end_at, created_at").
		From(taskTable)

	if len(filter.Ids) > 0 {
		query = query.Where(sq.Eq{"id": filter.Ids})
	}
	if len(filter.SatelliteIds) > 0 {
		query = query.Where(sq.Eq{"satellite_id": filter.SatelliteIds})
	}
	if len(filter.DeviceId) > 0 {
		query = query.Where(sq.Eq{"device_id": filter.DeviceId})
	}
	if len(filter.SignalType) > 0 {
		query = query.Where(sq.Eq{"signal_type": filter.SignalType})
	}
	if len(filter.GroupingType) > 0 {
		query = query.Where(sq.Eq{"grouping_type": filter.GroupingType})
	}
	if filter.StartAt != nil {
		query = query.Where(sq.GtOrEq{"start_at": *filter.StartAt})
	}
	if filter.EndAt != nil {
		query = query.Where(sq.LtOrEq{"end_at": *filter.EndAt})
	}
	if filter.Paginator.Page != 0 {
		query = query.Offset(filter.Paginator.Page)
	}
	if filter.Paginator.PerPage != 0 {
		query = query.Limit(filter.Paginator.PerPage)
	}

	var tasks []*model.Task
	if err := g.storage.db.Selectx(ctx, &tasks, query); err != nil {
		return nil, postgresError(err)
	}

	return tasks, nil
}

type ListSatellitesFilter struct {
	Ids                  []string
	ExternalSatelliteIds []string
	SatelliteName        []string
	Paginator            model.Paginator
}

func (g *GnssStore) ListSatellites(ctx context.Context, filter ListSatellitesFilter) ([]*model.SatelliteInfo, error) {
	query := g.storage.Builder().
		Select("id, external_satellite_id, satellite_name, created_at").
		From(satelliteTable)

	if len(filter.Ids) > 0 {
		query = query.Where(sq.Eq{"id": filter.Ids})
	}
	if len(filter.ExternalSatelliteIds) > 0 {
		query = query.Where(sq.Eq{"external_satellite_id": filter.ExternalSatelliteIds})
	}
	if len(filter.SatelliteName) > 0 {
		query = query.Where(sq.Eq{"satellite_name": filter.SatelliteName})
	}

	if filter.Paginator.Page != 0 {
		query = query.Offset(filter.Paginator.Page)
	}
	if filter.Paginator.PerPage != 0 {
		query = query.Limit(filter.Paginator.PerPage)
	}

	var satelliteInfo []*model.SatelliteInfo
	if err := g.storage.db.Selectx(ctx, &satelliteInfo, query); err != nil {
		return nil, postgresError(err)
	}

	return satelliteInfo, nil
}

type CreateSatelliteParams struct {
	ExternalSatelliteId string
	SatelliteName       string
}

func (g *GnssStore) CreateSatellite(ctx context.Context, params CreateSatelliteParams) (*model.SatelliteInfo, error) {
	query := g.storage.Builder().
		Insert(satelliteTable).
		SetMap(map[string]any{
			"external_satellite_id": params.ExternalSatelliteId,
			"satellite_name":        params.SatelliteName,
		}).
		Suffix("RETURNING id, external_satellite_id, satellite_name, created_at")

	var satelliteInfo model.SatelliteInfo
	if err := g.storage.db.Getx(ctx, &satelliteInfo, query); err != nil {
		return nil, postgresError(err)
	}

	return &satelliteInfo, nil
}

type DeleteDeviceFilter struct {
	Id string
}

func (g *GnssStore) DeleteDevice(ctx context.Context, filter DeleteDeviceFilter) error {
	query := g.storage.Builder().
		Delete(deviceTable).
		Where(sq.Eq{"id": filter.Id})

	if filter.Id != "" {
		query = query.Where(sq.Eq{"id": filter.Id})
	}

	if _, err := g.storage.db.Execx(ctx, query); err != nil {
		return postgresError(err)
	}

	return nil
}

func (g *GnssStore) AddSpectrum(ctx context.Context, req model.SpectrumRequest) error {
	query := g.storage.Builder().
		Insert("measurements_spectrum").
		SetMap(map[string]any{
			"spectrum":   req.Data.Spectrum,
			"start_freq": req.Data.StartFreq,
			"freq_step":  req.Data.FreqStep,
			"started_at": req.Data.StartTime,
		}).Suffix("RETURNING id")

	var id string
	if err := g.storage.db.Getx(ctx, &id, query); err != nil {
		return postgresError(err)
	}

	err := g.addHardwareMeasurement(ctx, req.Description, "", id, req.Token, "spectrum")
	if err != nil {
		return fmt.Errorf("failed to add hardware measurement spectrum: %w", err)
	}

	return nil
}

func (g *GnssStore) AddPower(ctx context.Context, req model.PowerRequest) error {
	query := g.storage.Builder().
		Insert("measurements_power").
		SetMap(map[string]any{
			"power":      req.Data.Power,
			"started_at": req.Data.StartTime,
			"time_step":  req.Data.TimeStep,
		}).Suffix("RETURNING id")

	var id string
	if err := g.storage.db.Getx(ctx, &id, query); err != nil {
		return postgresError(err)
	}

	err := g.addHardwareMeasurement(ctx, req.Description, id, "", req.Token, "power")
	if err != nil {
		return fmt.Errorf("failed to add hardware measurement power: %w", err)
	}
	return nil
}

func (g *GnssStore) AddPairMeasurement(ctx context.Context, req model.PairMeasurementRequest) error {
	queryPower := g.storage.Builder().
		Insert("measurements_power").
		SetMap(map[string]any{
			"power":      req.Power.Power,
			"started_at": req.Power.StartTime,
			"time_step":  req.Power.TimeStep,
		}).Suffix("RETURNING id")

	var powerId string
	if err := g.storage.db.Getx(ctx, &powerId, queryPower); err != nil {
		return postgresError(err)
	}

	querySpectrum := g.storage.Builder().
		Insert("measurements_spectrum").
		SetMap(map[string]any{
			"spectrum":   req.Spectrum.Spectrum,
			"start_freq": req.Spectrum.StartFreq,
			"freq_step":  req.Spectrum.FreqStep,
			"started_at": req.Spectrum.StartTime,
		}).Suffix("RETURNING id")

	var spectrumId string
	if err := g.storage.db.Getx(ctx, &spectrumId, querySpectrum); err != nil {
		return postgresError(err)
	}

	err := g.addHardwareMeasurement(ctx, req.Description, powerId, spectrumId, req.Token, "both")
	if err != nil {
		return fmt.Errorf("failed to add hardware measurement pair: %w", err)
	}
	return nil
}

func (g *GnssStore) addHardwareMeasurement(ctx context.Context, desc model.Description, measurementPowerID string, measurementSpectrumID string, token string, queryType string) error {
	switch queryType {
	case "spectrum":
		query := g.storage.Builder().
			Insert("hardware_measurements").
			SetMap(map[string]any{
				"token":                   token,
				"start_at":                desc.StartTime,
				"end_at":                  desc.EndTime,
				"group_type":              desc.Group,
				"signal":                  desc.Signal,
				"satellite_name":          desc.Target,
				"measurement_spectrum_id": measurementSpectrumID,
			})
		if _, err := g.storage.db.Execx(ctx, query); err != nil {
			return postgresError(err)
		}
	case "power":
		query := g.storage.Builder().
			Insert("hardware_measurements").
			SetMap(map[string]any{
				"token":                token,
				"start_at":             desc.StartTime,
				"end_at":               desc.EndTime,
				"group_type":           desc.Group,
				"signal":               desc.Signal,
				"satellite_name":       desc.Target,
				"measurement_power_id": measurementPowerID,
			})
		if _, err := g.storage.db.Execx(ctx, query); err != nil {
			return postgresError(err)
		}
	case "both":
		query := g.storage.Builder().
			Insert("hardware_measurements").
			SetMap(map[string]any{
				"token":                   token,
				"start_at":                desc.StartTime,
				"end_at":                  desc.EndTime,
				"group_type":              desc.Group,
				"signal":                  desc.Signal,
				"satellite_name":          desc.Target,
				"measurement_spectrum_id": measurementSpectrumID,
				"measurement_power_id":    measurementPowerID,
			})
		if _, err := g.storage.db.Execx(ctx, query); err != nil {
			return postgresError(err)
		}

	}
	return nil
}

func (g *GnssStore) CompareDeviceToken(ctx context.Context, token string) error {
	query := g.storage.Builder().
		Select("token").
		From("devices").
		Where(sq.Eq{"token": token})

	if _, err := g.storage.db.Execx(ctx, query); err != nil {
		return postgresError(err)
	}

	return nil
}

func (g *GnssStore) ListMeasurements(ctx context.Context, measurementReq model.MeasurementsFilter) ([]*model.Measurement, error) {
	query := g.storage.Builder().
		Select("hm.id", "hm.token", "hm.start_at", "hm.end_at", "hm.group_type", "hm.signal", "hm.satellite_name", "hm.measurement_power_id", "hm.measurement_spectrum_id").
		From("hardware_measurements hm")

	if measurementReq.Signal != nil {
		query = query.Where(sq.Eq{"hm.signal": *measurementReq.Signal})
	}
	if measurementReq.Group != nil {
		query = query.Where(sq.Eq{"hm.group_type": *measurementReq.Group})
	}
	if measurementReq.Target != nil {
		query = query.Where(sq.Eq{"hm.satellite_name": *measurementReq.Target})
	}
	if measurementReq.StartAt != nil {
		query = query.Where(sq.GtOrEq{"hm.start_at": *measurementReq.StartAt})
	}
	if measurementReq.EndAt != nil {
		query = query.Where(sq.LtOrEq{"hm.end_at": *measurementReq.EndAt})
	}
	if measurementReq.Token != nil {
		query = query.Where(sq.Eq{"hm.token": *measurementReq.Token})
	}

	var hardwareMeasurements []struct {
		Id            string    `db:"id"`
		Token         string    `db:"token"`
		StartAt       time.Time `db:"start_at"`
		EndAt         time.Time `db:"end_at"`
		GroupType     string    `db:"group_type"`
		Signal        string    `db:"signal"`
		SatelliteName string    `db:"satellite_name"`
		// MeasurementPowerID    string    `db:"measurement_power_id"`
		MeasurementSpectrumID string `db:"measurement_spectrum_id"`
	}
	if err := g.storage.db.Selectx(ctx, &hardwareMeasurements, query); err != nil {
		fmt.Println(hardwareMeasurements)
		return nil, postgresError(err)
	}

	var measurements []*model.Measurement

	for _, hm := range hardwareMeasurements {
		var measurement model.Measurement
		measurement.ID = hm.Id
		measurement.Token = hm.Token
		measurement.StartTime = hm.StartAt
		measurement.EndTime = hm.EndAt
		measurement.Group = hm.GroupType
		measurement.SignalType = hm.Signal
		measurement.Target = hm.SatelliteName

		// var powerDataDb struct {
		// 	Power     []float64 `db:"power"`
		// 	StartTime time.Time `db:"started_at"`
		// 	TimeStep  time.Time `db:"time_step"`
		// }
		// err := g.storage.db.Getx(ctx, &powerDataDb, g.storage.Builder().
		// 	Select("power", "started_at", "time_step").
		// 	From("measurements_power").
		// 	Where(sq.Eq{"id": hm.MeasurementPowerID}))
		// if err == nil {
		// 	measurement.DataPower = &model.DataPower{
		// 		Power:     powerDataDb.Power,
		// 		StartTime: powerDataDb.StartTime,
		// 		TimeStep:  powerDataDb.TimeStep,
		// 	}
		// } else if !(strings.Contains(err.Error(), pgx.ErrNoRows.Error()) || errors.Is(err, pgx.ErrNoRows)) {
		// 	return nil, postgresError(err)
		// }

		var spectrumDataDb struct {
			Spectrum  []float64 `db:"spectrum"`
			StartFreq float64   `db:"start_freq"`
			FreqStep  float64   `db:"freq_step"`
			StartedAt time.Time `db:"started_at"`
		}
		err := g.storage.db.Getx(ctx, &spectrumDataDb, g.storage.Builder().
			Select("spectrum", "start_freq", "freq_step", "started_at").
			From("measurements_spectrum").
			Where(sq.Eq{"id": hm.MeasurementSpectrumID}))
		if err == nil {
			measurement.DataSpectrum = &model.DataSpectrum{
				Spectrum:  spectrumDataDb.Spectrum,
				StartFreq: spectrumDataDb.StartFreq,
				FreqStep:  spectrumDataDb.FreqStep,
				StartTime: spectrumDataDb.StartedAt,
			}
		} else if !(strings.Contains(err.Error(), pgx.ErrNoRows.Error()) || errors.Is(err, pgx.ErrNoRows)) {
			return nil, postgresError(err)
		}

		measurements = append(measurements, &measurement)
	}

	return measurements, nil
}

func (g *GnssStore) RinexList(ctx context.Context) ([]*model.RinexResults, error) {
	jsonData := []*model.RinexResults{
		{
			Header: &model.Header{
				RinexVersion:       "2.11",
				FileType:           "OBSERVATION DATA",
				PgmRunByDate:       "CONVBIN 2.4.3 20241106 095923 UTC",
				Comments:           []string{"format: Javad GREIS", "log: ../tmp_rinex2_hourly/YARO_2024_11_06_08.jps"},
				MarkerName:         "",
				MarkerNumber:       "",
				ObserverAgency:     "",
				RecInfo:            "",
				AntInfo:            "",
				ApproxPositionXyz:  []float64{2626719.2512, 2195458.1382, 5363666.0981},
				AntennaDeltaHen:    []float64{0.0000, 0.0000, 0.0000},
				WavelengthFactL1L2: []int{1, 1},
				TypesOfObs:         []string{"C1", "L1", "P1", "P2", "L2", "C2", "C5", "L5"},
				Interval:           10.000,
				TimeOfFirstObs:     "2024 11 06 08 59 50.0000000 GPS",
				TimeOfLastObs:      "2024 11 06 08 59 50.0000000 GPS",
				EndOfHeader:        true,
			},
			Observations: []*model.Observation{
				{
					Time:      "2024 11 06 08 59 50.0000000",
					EpochFlag: 0,
					Satellites: []*model.Satellite{
						{
							SatelliteID:  "23R",
							Observations: []string{"21637955.781", "115748497.2511", "21637957.379"},
						},
						{
							SatelliteID:  "23R",
							Observations: []string{"19723101.973", "105246184.4061", "19723098.435"},
						},
						{
							SatelliteID:  "22G",
							Observations: []string{"22766526.105", "121529156.8331", "22766524.885", "22766542.459", "94522675.6331", "22766539.847"},
						},
						{
							SatelliteID:  "16G",
							Observations: []string{"22303210.999", "117204258.0021", "22303209.896", "22303216.932", "91328033.4211"},
						},
						{
							SatelliteID:  "11G",
							Observations: []string{"23210320.069", "121971045.1601", "23210320.441", "23210322.845", "95042373.3311", "23210322.281", "23210330.328", "91082273.6561"},
						},
						{
							SatelliteID:  "06S",
							Observations: []string{"22580005.523", "118658672.7071", "22580005.781", "22580013.192", "92461282.6061", "22580012.928", "22580016.610", "88608732.2241"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"42698660.033", "224383954.4171", "42698671.734", "167558951.1841"},
						},
						{
							SatelliteID:  "31R",
							Observations: []string{"23697554.619", "124531426.1341", "23697555.768", "23697562.681", "97037457.3321", "23697564.734"},
						},
						{
							SatelliteID:  "09G",
							Observations: []string{"21704140.137", "115898995.3911", "21704139.952", "21704155.235", "90143723.5771", "21704154.731"},
						},
						{
							SatelliteID:  "04S",
							Observations: []string{"20198608.366", "106144377.4681", "20198607.665", "20198612.965", "82709895.0251", "20198613.097", "20198617.666", "79263647.6451"},
						},
						{
							SatelliteID:  "27G",
							Observations: []string{"42093623.760", "221203304.4561", "42093645.171", "165184229.9381"},
						},
						{
							SatelliteID:  "07G",
							Observations: []string{"23117956.159", "121485787.6041", "23117956.168", "23117965.291", "94664269.7691", "23117965.423"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"43006100.412", "225998406.6231", "43006124.803", "168764996.4521"},
						},
						{
							SatelliteID:  "17G",
							Observations: []string{"23183933.938", "124148883.1081", "23183929.689", "23183941.609", "96560237.9341", "23183941.843"},
						},
						{
							SatelliteID:  "11G",
							Observations: []string{"22758365.505", "119595850.0051", "22758364.780", "22758377.695", "93191533.4211", "22758378.609", "22758382.941", "89308544.0321"},
						},
						{
							SatelliteID:  "06S",
							Observations: []string{"42183829.071", "221677774.9391"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"24446368.591", "128466528.5021", "24446366.328", "24446370.693", "100103783.4071"},
						},
						{
							SatelliteID:  "21G",
							Observations: []string{"22186059.844", "118555649.8941", "22186058.975", "22186070.877", "92209976.1571", "22186070.898"},
						},
						{
							SatelliteID:  "17G",
							Observations: []string{"20159372.252", "105938230.3921", "20159371.667", "20159377.217", "82549266.1111", "20159377.804", "20159379.855", "79109712.9511"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"43094771.706", "226463949.3301"},
						},
						{
							SatelliteID:  "19G",
							Observations: []string{"19609198.442", "104969558.3841", "19609199.252", "19609208.569", "81642989.0881", "19609208.989"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"19833957.468", "105949631.2851", "19833957.537", "19833967.769", "82405321.5061", "19833967.832"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"22492129.204", "118197052.4691", "22492128.520", "22492138.692", "92101631.6171", "22492137.919", "22492142.173", "88264066.4751"},
						},
						{
							SatelliteID:  "23R",
							Observations: []string{"23243078.307", "124334761.901", "23243080.966"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"21633253.668", "115439046.769", "21633253.278"},
						},
						{
							SatelliteID:  "16G",
							Observations: []string{"23998309.120", "128419771.441", "23998309.072", "23998324.698", "99882054.881", "23998324.952"},
						},
						{
							SatelliteID:  "11G",
							Observations: []string{"21923664.733", "115209766.128", "21923663.791", "21923673.750", "89773889.039"},
						},
						{
							SatelliteID:  "16G",
							Observations: []string{"23424519.852", "123096630.383", "23424516.884", "23424524.127", "95919443.589", "23424522.428", "23424531.355", "91922797.244"},
						},
						{
							SatelliteID:  "23R",
							Observations: []string{"24453451.191", "128503533.158", "24453450.555", "24453466.978", "100132562.228", "24453469.047", "24453473.966", "95960366.409"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"42736359.492", "224582024.345", "42736372.074", "167706855.657"},
						},
						{
							SatelliteID:  "09G",
							Observations: []string{"19674729.455", "105062138.974", "19674728.715", "19674739.645", "81715078.893", "19674740.188"},
						},
						{
							SatelliteID:  "23G",
							Observations: []string{"21459301.640", "112769325.037", "21459302.056", "21459312.855", "87872178.458", "21459310.843", "21459317.637", "84210833.102"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"42140500.115", "221449640.742", "42140522.734", "165368224.536"},
						},
						{
							SatelliteID:  "23G",
							Observations: []string{"21459110.045", "112768551.155", "21459110.399", "21459115.217", "87871632.597", "21459115.391"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"23600946.305", "124023891.606", "23600945.882", "23600958.488", "96642021.833", "23600957.607", "23600963.687", "92615276.176"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"43044614.114", "226200813.895", "43044639.224", "168916147.294"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"20826002.217", "111522298.111", "20826001.503", "20826012.125", "86739563.905", "20826011.753"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"22444434.682", "120020508.181", "22444435.129", "22444449.165", "93349281.012", "22444449.372"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"42219944.076", "221867534.076"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"22459782.002", "118026955.097", "22459779.733", "22459789.012", "91969049.478"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"20067559.460", "105455729.757", "20067558.644", "20067565.306", "82173286.724", "20067565.276", "20067568.705", "78749398.299"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"43131931.362", "226659202.146"},
						},
						{
							SatelliteID:  "19G",
							Observations: []string{"19132570.585", "102418109.420", "19132570.750", "19132580.268", "79658524.723", "19132580.892"},
						},
						{
							SatelliteID:  "26G",
							Observations: []string{"22370985.665", "119250026.543", "22370984.373"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"20677759.178", "110457042.088", "20677758.516", "20677771.362", "85911078.807", "20677771.886"},
						},
						{
							SatelliteID:  "26G",
							Observations: []string{"23932031.753", "125763737.830", "23932031.816", "23932044.647", "97997735.733", "23932045.849", "23932049.435", "93914496.482"},
						},
					},
				},
			},
		},
	}

	return jsonData, nil
}
