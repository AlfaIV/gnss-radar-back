package store

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	gnssTable   = "gnss_coords"
	deviceTable = "devices"
)

type IGnssStore interface {
	ListGnssCoords(ctx context.Context, filter ListGnssCoordsFilter) ([]*model.GnssCoords, error)
	UpsetDevice(ctx context.Context, params UpsetDeviceParams) (*model.Device, error)
	ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error)
}

type GnssStore struct {
	storage *Storage
}

func (g *GnssStore) ListDevice(ctx context.Context, filter ListDeviceFilter) ([]*model.Device, error) {
	query := g.storage.Builder().
		Select("id, name, token, description, x, y, z").
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
		Select("id, satellite_id, satellite_name, x, y, z").
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

type UpsetDeviceParams struct {
	Name        string  `db:"name"`
	Token       string  `db:"token"`
	Description *string `db:"description"`
	Coords      model.Coords
}

func (g *GnssStore) UpsetDevice(ctx context.Context, params UpsetDeviceParams) (*model.Device, error) {
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
		Suffix(`                                                            
			ON CONFLICT (name) DO UPDATE SET
			token = EXCLUDED.token,
			description = EXCLUDED.description,
			x = EXCLUDED.x,
			y = EXCLUDED.y,
			z = EXCLUDED.z
		`).
		Suffix("RETURNING id, name, token, description, x, y, z")

	var device *model.Device
	if err := g.storage.db.Selectx(ctx, &device, query); err != nil {
		return nil, postgresError(err)
	}

	return device, nil
}

type ListDeviceFilter struct {
	Ids       []string `db:"id"`
	Names     []string `db:"name"`
	Tokens    []string `db:"token"`
	Paginator model.Paginator
}
