package service

import (
	"context"
	"fmt"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"strconv"

	"github.com/Gokert/gnss-radar/internal/store"
)

type IGnss interface {
	ListGnss(ctx context.Context, req ListGnssRequest) ([]*model.GnssCoords, error)
	ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error)
	UpsetDevice(ctx context.Context, params UpsetDeviceParams) (*model.Device, error)
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
		return nil, fmt.Errorf("gnssStore.List: %w", err)
	}

	return gnss, nil
}

type UpsetDeviceParams struct {
	Name        string
	Token       string
	Description *string
	X           string
	Y           string
	Z           string
}

func (g *GnssService) UpsetDevice(ctx context.Context, params UpsetDeviceParams) (*model.Device, error) {
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

	device, err := g.gnssStore.UpsetDevice(ctx, store.UpsetDeviceParams{
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
		return nil, fmt.Errorf("gnssStore.UpsetDevice: %w", err)
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
