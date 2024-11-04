package store

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
)

const (
	gnssTable = "gnss"
)

type IGnssStore interface {
	List(ctx context.Context, params ListParams) ([]*model.Gnss, error)
}

type GnssStore struct {
}

func NewGnssStore() *GnssStore {
	return &GnssStore{}
}

type ListParams struct {
	X float64
	Y float64
	Z float64
}

// func (g *GnssStore) List(ctx context.Context, params ListParams) ([]*model.Gnss, error) {
// 	query := g.storage.Builder().
// 		Select("name, x, y, z").From(gnssTable)

// 	if params.X != 0 {
// 		query = query.Where("x = ?", params.X)
// 	}
// 	if params.Y != 0 {
// 		query = query.Where("y = ?", params.Y)
// 	}
// 	if params.Z != 0 {
// 		query = query.Where("z = ?", params.Z)
// 	}

// 	var coords []*model.Coord
// 	err := g.storage.Query(query, &coords)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return coords, nil
// }

func (g *GnssStore) List(ctx context.Context, params ListParams) ([]*model.Gnss, error) {
	jsonData := []*model.Gnss{
		{
			ID:            "PC06",
			SatelliteID:   "PC06",
			SatelliteName: "PC06",
			Coordinates: &model.CoordsResults{
				X: "-16806.320344",
				Y: "29291.120310",
				Z: "-25355.710938",
			},
		},
		{
			ID:            "PC07",
			SatelliteID:   "PC07",
			SatelliteName: "PC07",
			Coordinates: &model.CoordsResults{
				X: "-6959.418476",
				Y: "39332.954409",
				Z: "-13000.851001",
			},
		},
		{
			ID:            "PC08",
			SatelliteID:   "PC08",
			SatelliteName: "PC08",
			Coordinates: &model.CoordsResults{
				X: "-1908.204600",
				Y: "21553.224987",
				Z: "36203.881809",
			},
		},
		{
			ID:            "PC09",
			SatelliteID:   "PC09",
			SatelliteName: "PC09",
			Coordinates: &model.CoordsResults{
				X: "-11202.586298",
				Y: "28046.331947",
				Z: "-29182.143554",
			},
		},
		{
			ID:            "PC10",
			SatelliteID:   "PC10",
			SatelliteName: "PC10",
			Coordinates: &model.CoordsResults{
				X: "-917.431406",
				Y: "41238.966109",
				Z: "-6711.991412",
			},
		},
		{
			ID:            "PC11",
			SatelliteID:   "PC11",
			SatelliteName: "PC11",
			Coordinates: &model.CoordsResults{
				X: "-16138.177056",
				Y: "-3913.891460",
				Z: "-22348.411693",
			},
		},
		{
			ID:            "PC12",
			SatelliteID:   "PC12",
			SatelliteName: "PC12",
			Coordinates: &model.CoordsResults{
				X: "-997.099233",
				Y: "-19759.345910",
				Z: "-19638.934483",
			},
		},
		{
			ID:            "PC13",
			SatelliteID:   "PC13",
			SatelliteName: "PC13",
			Coordinates: &model.CoordsResults{
				X: "5858.392549",
				Y: "25505.986419",
				Z: "33308.911170",
			},
		},
		{
			ID:            "PC14",
			SatelliteID:   "PC14",
			SatelliteName: "PC14",
			Coordinates: &model.CoordsResults{
				X: "-17706.605729",
				Y: "-14691.268566",
				Z: "15829.680477",
			},
		},
		{
			ID:            "PC16",
			SatelliteID:   "PC16",
			SatelliteName: "PC16",
			Coordinates: &model.CoordsResults{
				X: "-22387.055407",
				Y: "28560.640995",
				Z: "-21454.026667",
			},
		},
	}

	return jsonData, nil
}
