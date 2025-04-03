package measurements_domain_gateway

import (
	"context"
	"io"
)

type EphemerisFileMeta struct {
	Name     string `json:"name"`
	Datetime string `json:"datetime"`
}

type GetEphemerisResponse struct {
	Ephemeris []EphemerisFileMeta `json:"ephemeris"`
	Total     uint64              `json:"total"`
	Page      uint64              `json:"page"`
}

type SatelliteData struct {
	Group     string  `json:"group"`
	Name      string  `json:"name"`
	Azimuth   float64 `json:"azimuth"`
	Range     float64 `json:"range"`
	Elevation float64 `json:"elevation"`
}

type Satellites struct {
	Satellites []SatelliteData `json:"azimuth"`
}

type Usecase interface {
	GetEphemeris(ctx context.Context, page uint64, size uint64) ([]EphemerisFileMeta, uint64, error)
	LoadEphemeris(ctx context.Context, file io.Reader, name string) error
	GetSatellitesPosition(ctx context.Context) (Satellites, error)
}
