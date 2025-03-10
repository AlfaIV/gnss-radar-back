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

type Usecase interface {
	GetEphemeris(ctx context.Context, page uint64, size uint64) ([]EphemerisFileMeta, uint64, error)
	LoadEphemeris(ctx context.Context, file io.Reader, name string) error
}
