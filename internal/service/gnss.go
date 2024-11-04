package service

import (
	"context"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
)

type ListRequest struct {
	X string
	Y string
	Z string
}

type IGnss interface {
	ListGnss(ctx context.Context, req ListRequest) (*model.GNSSPagination, error)
}

type GnssService struct {
	store store.IGnssStore
}

func NewGnssService(store store.IGnssStore) *GnssService {
	return &GnssService{store: store}
}

func (g *GnssService) ListGnss(ctx context.Context, req ListRequest) (*model.GNSSPagination, error) {
	// Xf, err := strconv.ParseFloat(req.X, 64)
	// if err != nil {
	// 	return nil, err
	// }
	// yf, err := strconv.ParseFloat(req.Y, 64)
	// if err != nil {
	// 	return nil, err
	// }
	// zf, err := strconv.ParseFloat(req.Z, 64)
	// if err != nil {
	// 	return nil, err
	// }

	// coords := store.ListParams{X: Xf, Y: yf, Z: zf}
	gnss, err := g.store.List()
	if err != nil {
		return nil, err
	}

	var ans *model.GNSSPagination
	ans.Items = gnss

	return ans, nil
}
