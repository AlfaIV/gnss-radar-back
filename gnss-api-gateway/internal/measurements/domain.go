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
	Group     string
	Name      string
	Azumuth   float64
	Range     float64
	Elevation float64
}

type Satellites struct {
	Satellites []SatelliteData
}

type Usecase interface {
	GetEphemeris(ctx context.Context, page uint64, size uint64) ([]EphemerisFileMeta, uint64, error)
	LoadEphemeris(ctx context.Context, file io.Reader, name string) error
	GetSatellitesPosition(ctx context.Context) (Satellites, error)
}

//TAH4UK

type SatelliteInterval struct {
	StartDatetime string
	EndDatetime   string
}

type SatelliteWithIntervals struct {
	Group     string
	Name      string
	Intervals []SatelliteInterval
}
