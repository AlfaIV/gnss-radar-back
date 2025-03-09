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

type Repository interface {
	UploadEphemeris(ctx context.Context, req EphemerisToLoad) error
	GetEphemeris(ctx context.Context, req PaginatedRequest) ([]EphemerisFileMeta, error)
}
