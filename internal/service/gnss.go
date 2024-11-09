package service

import (
	"context"
	"fmt"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
	"strconv"
)

type IGnss interface {
	ListGnss(ctx context.Context, req ListGnssRequest) ([]*model.GnssCoords, error)
	ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error)
	CreateDevice(ctx context.Context, params CreateDeviceParams) (*model.Device, error)
	UpdateDevice(ctx context.Context, params UpdateDeviceParams) (*model.Device, error)
	RinexList(ctx context.Context, req RinexRequest) ([]*model.RinexResults, error)
	CreateTask(ctx context.Context, params store.CreateTaskParams) (*model.Task, error)
	UpdateTask(ctx context.Context, params store.UpdateTaskParams) (*model.Task, error)
	DeleteTask(ctx context.Context, filter store.DeleteTaskFilter) error
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
	Token       string
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

	device, err := g.gnssStore.CreateDevice(ctx, store.CreateDeviceParams{
		Name:        params.Name,
		Token:       params.Token,
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
	Token       string
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
		Token:       params.Token,
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
	task, err := g.gnssStore.CreateTask(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.CreateTask: %w", err)
	}

	return task, nil
}

func (g *GnssService) UpdateTask(ctx context.Context, params store.UpdateTaskParams) (*model.Task, error) {
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
