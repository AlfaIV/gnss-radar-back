package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/samber/lo"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
	"github.com/google/uuid"
)

type IGnss interface {
	ListGnss(ctx context.Context, req ListGnssRequest) ([]*model.GnssCoords, error)
	ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error)
	CreateDevice(ctx context.Context, params CreateDeviceParams) (*model.Device, error)
	UpdateDevice(ctx context.Context, params UpdateDeviceParams) (*model.Device, error)
	DeleteDevice(ctx context.Context, filter store.DeleteDeviceFilter) error
	RinexList(ctx context.Context, req RinexRequest) ([]*model.RinexResults, error)
	CreateTask(ctx context.Context, params store.CreateTaskParams) (*model.Task, error)
	UpdateTask(ctx context.Context, params store.UpdateTaskParams) (*model.Task, error)
	DeleteTask(ctx context.Context, filter store.DeleteTaskFilter) error
	ListTasks(ctx context.Context, filter ListTasksFilter) ([]*model.Task, error)
	ListSatellites(ctx context.Context, filter ListSatellitesFilter) ([]*model.SatelliteInfo, error)
	CreateSatellite(ctx context.Context, params store.CreateSatelliteParams) (*model.SatelliteInfo, error)
	ListMeasurements(ctx context.Context, measurementReq model.MeasurementsFilter) ([]*model.Measurement, error)
}

type GnssService struct {
	gnssStore store.IGnssStore
}

func NewGnssService(store store.IGnssStore) *GnssService {
	return &GnssService{gnssStore: store}
}

type ListGnssRequest struct {
	X         string
	Y         string
	Z         string
	Paginator model.Paginator
}

func (g *GnssService) ListGnss(ctx context.Context, req ListGnssRequest) ([]*model.GnssCoords, error) {
	xf, err := strconv.ParseFloat(req.X, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	yf, err := strconv.ParseFloat(req.Y, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	zf, err := strconv.ParseFloat(req.Z, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}

	gnss, err := g.gnssStore.ListGnssCoords(ctx, store.ListGnssCoordsFilter{
		X: xf, Y: yf, Z: zf,
		Paginator: req.Paginator,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListGnssCoords: %w", err)
	}

	return gnss, nil
}

type CreateDeviceParams struct {
	Name        string
	Description *string
	X           string
	Y           string
	Z           string
}

func (g *GnssService) CreateDevice(ctx context.Context, params CreateDeviceParams) (*model.Device, error) {
	xf, err := strconv.ParseFloat(params.X, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	yf, err := strconv.ParseFloat(params.Y, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	zf, err := strconv.ParseFloat(params.Z, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}

	token := uuid.NewString()

	device, err := g.gnssStore.CreateDevice(ctx, store.CreateDeviceParams{
		Name:        params.Name,
		Token:       token,
		Description: params.Description,
		Coords: model.Coords{
			X: xf,
			Y: yf,
			Z: zf,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.CreateDevice: %w", err)
	}

	return device, nil
}

type UpdateDeviceParams struct {
	ID          string
	Name        string
	Description *string
	X           string
	Y           string
	Z           string
}

func (g *GnssService) UpdateDevice(ctx context.Context, params UpdateDeviceParams) (*model.Device, error) {
	xf, err := strconv.ParseFloat(params.X, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	yf, err := strconv.ParseFloat(params.Y, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	zf, err := strconv.ParseFloat(params.Z, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}

	device, err := g.gnssStore.UpdateDevice(ctx, store.UpdateDeviceParams{
		Id:          params.ID,
		Name:        params.Name,
		Description: params.Description,
		Coords: model.Coords{
			X: xf,
			Y: yf,
			Z: zf,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.UpdateDevice: %w", err)
	}

	return device, nil
}

type ListDeviceFilter struct {
	Ids       []string
	Names     []string
	Tokens    []string
	Paginator model.Paginator
}

func (g *GnssService) ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error) {
	device, err := g.gnssStore.ListDevice(ctx, store.ListDeviceFilter{
		Ids: filter.Ids, Names: filter.Names, Tokens: filter.Tokens,
		Paginator: filter.Paginator,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListDevice: %w", err)
	}

	return device, nil
}

type RinexRequest struct{}

func (g *GnssService) RinexList(ctx context.Context, req RinexRequest) ([]*model.RinexResults, error) {
	rinexlist, err := g.gnssStore.RinexList(ctx)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListRinexlist: %w", err)
	}

	return rinexlist, nil
}

func (g *GnssService) CreateTask(ctx context.Context, params store.CreateTaskParams) (*model.Task, error) {
	listTask, err := g.gnssStore.ListTask(ctx, store.ListTasksFilter{
		DeviceId: []string{params.DeviceId},
		StartAt:  &params.StartAt,
		EndAt:    &params.EndAt,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListTask: %w", err)
	}

	if len(listTask) > 0 {
		return nil, store.ErrEntityAlreadyExist
	}

	satellites, err := g.gnssStore.ListSatellites(ctx, store.ListSatellitesFilter{Ids: []string{params.SatelliteId}})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListSatellites: %w", err)
	}
	if len(satellites) == 0 {
		return nil, fmt.Errorf("satellites with id = %s not found", params.SatelliteId)
	}

	devices, err := g.gnssStore.ListDevice(ctx, store.ListDeviceFilter{Ids: []string{params.DeviceId}})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListDevice: %w", err)
	}
	if len(devices) == 0 {
		return nil, fmt.Errorf("devices with id = %s not found", params.DeviceId)
	}

	task, err := g.gnssStore.CreateTask(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.CreateTask: %w", err)
	}

	return task, nil
}

func (g *GnssService) UpdateTask(ctx context.Context, params store.UpdateTaskParams) (*model.Task, error) {
	listTask, err := g.gnssStore.ListTask(ctx, store.ListTasksFilter{
		DeviceId: []string{params.DeviceID},
		StartAt:  &params.StartAt,
		EndAt:    &params.EndAt,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListTask: %w", err)
	}

	if len(listTask) > 0 {
		return nil, store.ErrEntityAlreadyExist
	}

	satellites, err := g.gnssStore.ListSatellites(ctx, store.ListSatellitesFilter{Ids: []string{params.SatelliteID}})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListSatellites: %w", err)
	}
	if len(satellites) == 0 {
		return nil, fmt.Errorf("satellites with id = %s not found", params.SatelliteID)
	}

	task, err := g.gnssStore.UpdateTask(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.UpdateTask: %w", err)
	}

	return task, nil
}

func (g *GnssService) DeleteTask(ctx context.Context, filter store.DeleteTaskFilter) error {
	err := g.gnssStore.DeleteTask(ctx, filter)
	if err != nil {
		return fmt.Errorf("gnssStore.DeleteTask: %w", err)
	}

	return nil
}

type ListTasksFilter struct {
	Ids           []string
	SatelliteIds  []string
	SatelliteName []string
	SignalType    []model.SignalType
	GroupingType  []model.GroupingType
	StartAt       *time.Time
	EndAt         *time.Time
	Paginator     model.Paginator
}

func (g *GnssService) ListTasks(ctx context.Context, filter ListTasksFilter) ([]*model.Task, error) {
	var satelliteIds []string

	if len(filter.SatelliteName) > 0 {
		satellites, err := g.gnssStore.ListSatellites(ctx, store.ListSatellitesFilter{
			SatelliteName: filter.SatelliteName,
		})
		if err != nil {
			return nil, fmt.Errorf("gnssStore.ListSatellites: %w", err)
		}

		if len(satellites) == 0 {
			return []*model.Task{}, nil
		}

		satelliteIds = lo.Map(satellites, func(satellite *model.SatelliteInfo, _ int) string {
			return satellite.ID
		})
	}

	tasks, err := g.gnssStore.ListTask(ctx, store.ListTasksFilter{
		Ids:          append(satelliteIds, filter.Ids...),
		SatelliteIds: filter.SatelliteIds,
		SignalType:   filter.SignalType,
		GroupingType: filter.GroupingType,
		StartAt:      filter.StartAt,
		EndAt:        filter.EndAt,
		Paginator:    filter.Paginator,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListTask: %w", err)
	}

	return tasks, nil
}

type ListSatellitesFilter struct {
	Ids                  []string
	ExternalSatelliteIds []string
	SatelliteName        []string
	DeviceIds            []string
	Paginator            model.Paginator
}

func (g *GnssService) ListSatellites(ctx context.Context, filter ListSatellitesFilter) ([]*model.SatelliteInfo, error) {
	if len(filter.DeviceIds) > 0 {
		devices, err := g.gnssStore.ListDevice(ctx, store.ListDeviceFilter{
			Ids: filter.DeviceIds,
		})
		if err != nil {
			return nil, fmt.Errorf("gnssStore.ListDevice: %w", err)
		}

		if len(devices) == 0 {
			return []*model.SatelliteInfo{}, nil
		}

		filter.Ids = lo.Map(devices, func(device *model.Device, _ int) string {
			return device.ID
		})
	}

	satellites, err := g.gnssStore.ListSatellites(ctx, store.ListSatellitesFilter{
		Ids:                  filter.Ids,
		ExternalSatelliteIds: filter.ExternalSatelliteIds,
		Paginator:            filter.Paginator,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListSatellites: %w", err)
	}

	return satellites, nil
}

func (g *GnssService) CreateSatellite(ctx context.Context, params store.CreateSatelliteParams) (*model.SatelliteInfo, error) {
	satellite, err := g.gnssStore.CreateSatellite(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.CreateSatellite: %w", err)
	}

	return satellite, nil
}

func (g *GnssService) DeleteDevice(ctx context.Context, filter store.DeleteDeviceFilter) error {
	err := g.gnssStore.DeleteDevice(ctx, filter)
	if err != nil {
		return fmt.Errorf("gnssStore.DeleteDevice: %w", err)
	}

	return nil
}

func (g *GnssService) ListMeasurements(ctx context.Context, measurementReq model.MeasurementsFilter) ([]*model.Measurement, error) {
	measurements, err := g.gnssStore.ListMeasurements(ctx, measurementReq)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListMeasurements: %w", err)
	}

	return measurements, nil
}
