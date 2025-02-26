package statistics_repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type PgxIFace interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type UserRepo struct {
	pool   PgxIFace
	logger *logrus.Logger
}

func NewUserRepo(pool PgxIFace, logger *logrus.Logger) *UserRepo {
	return &UserRepo{pool: pool, logger: logger}
}