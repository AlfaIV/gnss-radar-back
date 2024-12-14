package executor

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}

// SqlizedExecer Exec с помощью Squirrel
type SqlizedExecer interface {
	Execx(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error)
}

// SqlizedQuerier Query с помощью Squirrel
type SqlizedQuerier interface {
	Queryx(ctx context.Context, sqlizer Sqlizer) (pgx.Rows, error)
	QueryxRow(ctx context.Context, sqlizer Sqlizer) pgx.Row
}

// SqlizedGetter Get с помощью Squirrel
type SqlizedGetter interface {
	Get(ctx context.Context, destPtr interface{}, sql string, args ...interface{}) error
	Getx(ctx context.Context, destPtr interface{}, sqlizer Sqlizer) error
}

// SqlizedSelecter Select с помощью Squirrel
type SqlizedSelecter interface {
	Select(ctx context.Context, destSlice interface{}, sql string, args ...interface{}) error
	Selectx(ctx context.Context, destSlice interface{}, sqlizer Sqlizer) error
}

// BasicExecutor - интерфейс с базовыми методами, на которых строятся операции интерфейса Executor.
type BasicExecutor interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// Executor - расширяет BasicExecutor методами для работы с Sqlizer.
type Executor interface {
	BasicExecutor
	SqlizedExecer
	SqlizedQuerier
	SqlizedGetter
	SqlizedSelecter
}

type ExecutorImpl struct {
	BasicExecutor
}

type errRow struct {
	error
}

func (e *ExecutorImpl) Execx(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return nil, err
	}
	return e.Exec(ctx, sql, args...)
}

func (e *ExecutorImpl) Queryx(ctx context.Context, sqlizer Sqlizer) (pgx.Rows, error) {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return nil, err
	}
	return e.Query(ctx, sql, args...)
}

func (e *ExecutorImpl) QueryxRow(ctx context.Context, sqlizer Sqlizer) pgx.Row {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return errRow{err}
	}
	return e.QueryRow(ctx, sql, args...)
}

func (er errRow) Scan(...interface{}) error {
	return er.error
}

func (e *ExecutorImpl) Getx(ctx context.Context, destPtr interface{}, sqlizer Sqlizer) error {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	return e.Get(ctx, destPtr, sql, args...)
}

func (e *ExecutorImpl) Select(ctx context.Context, destSlice interface{}, sql string, args ...interface{}) error {
	err := pgxscan.Select(ctx, e, destSlice, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	return err
}

func (e *ExecutorImpl) Selectx(ctx context.Context, destSlice interface{}, sqlizer Sqlizer) error {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	return e.Select(ctx, destSlice, sql, args...)
}

func NewExecutor(pool *pgxpool.Pool) *ExecutorImpl {
	return &ExecutorImpl{
		BasicExecutor: pool,
	}
}

func (e *ExecutorImpl) Get(ctx context.Context, destPtr interface{}, sql string, args ...interface{}) error {
	return pgxscan.Get(ctx, e, destPtr, sql, args...)
}
