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

type SatelliteData struct {
	Group string
	Name string
	Azumuth float64
	Range float64
	Elevation float64
}

type Satellites struct {
	Satellites []SatelliteData
}

type Repository interface {
	UploadEphemeris(ctx context.Context, req EphemerisToLoad) error
	GetEphemeris(ctx context.Context, req PaginatedRequest) ([]EphemerisFileMeta, uint64, error)
	GetSatellitesCoordinates(ctx context.Context) (Satellites, error)
}
