package measurements_domain_gateway

import (
	"context"
	"io"
)

type EphemerisFileMeta struct {
	Name     string	`json:"name"`
	Datetime string `json:"datetime"`
}

type Usecase interface {
	GetEphemeris(ctx context.Context, page uint64, size uint64) ([]EphemerisFileMeta, error)
	LoadEphemeris(ctx context.Context, file io.Reader, name string) error
}
