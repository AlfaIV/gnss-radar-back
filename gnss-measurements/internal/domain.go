package measurements_domain

import (
	"context"
	"io"
)

type EphemerisFileMeta struct {
	Name     string
	Datetime string
}

type PaginatedRequest struct {
	Page uint64
	Size uint64
}

type EphemerisToLoad struct {
	Name        string
	MinioName   string
	Payload     io.Reader
	PayloadSize int64
	ContentType string
}

type Satellites struct {
	Satellites []Satellite `json:"Satellites"`
}

type Satellite struct {
	Group     string  `json:"Group"`
	Name      string  `json:"Name"`
	Azimuth   float64 `json:"Azimuth"`
	Elevation float64 `json:"Elevation"`
	Range     float64 `json:"Range"`
}

type SatelliteInterval struct {
	StartDatime string `json:"startDatetime"`
	EndDatetime string `json:"endDatetime"`
}

type SatelliteWithIntervals struct {
	Group     string              `json:"group"`
	Name      string              `json:"name"`
	Intervals []SatelliteInterval `json:"intervals"`
}

type IntervalSatellitesResponse struct {
	Satellites []SatelliteWithIntervals
}

type Repository interface {
	UploadEphemeris(ctx context.Context, req EphemerisToLoad) error
	GetEphemeris(ctx context.Context, req PaginatedRequest) ([]EphemerisFileMeta, uint64, error)
	GetSatellitesCoordinates(ctx context.Context) (Satellites, error)
	GetSatellitesIntervals(ctx context.Context, startDatetime string, endDatetime string, satellites []string) ([]SatelliteWithIntervals, error)
}
