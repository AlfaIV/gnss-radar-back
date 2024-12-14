package store

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"strings"
)

var (
	// ErrNotFound Ошибка при отсутствии сущности
	ErrNotFound = errors.New("entity not found")
	// ErrLockedEntityAccess Ошибка при попытке получить доступ к заблокированной сущности
	ErrLockedEntityAccess = errors.New("an attempt to acquire an access to the locked entity")
	// ErrOneOfFilterRequired Ошибка требуется один из фильтров
	ErrOneOfFilterRequired = errors.New("one of filter is required")
	// ErrEntityAlreadyExist Ошибка при попытке создать сущность, которая уже существует
	ErrEntityAlreadyExist = errors.New("entity already exists")
)

const (
	// PgErrConcurrentLockAcquisition Код ошибки при попытке получить доступ к заблокированной сущности
	PgErrConcurrentLockAcquisition = "55P03"
	// PgErrCodeUniqueViolation Код ошибки при нарушении уникальности
	PgErrCodeUniqueViolation = "23505"
)

func postgresError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), pgx.ErrNoRows.Error()) || errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	var (
		code      string
		pgConnErr *pgconn.PgError
		pgErr     *pq.Error
	)
	if ok := errors.As(err, &pgConnErr); ok {
		code = pgConnErr.Code
	}
	if ok := errors.As(err, &pgErr); ok {
		code = string(pgErr.Code)
	}

	switch code {
	case PgErrCodeUniqueViolation:
		return fmt.Errorf("%w: %w", ErrEntityAlreadyExist, err)
	case PgErrConcurrentLockAcquisition:
		return ErrLockedEntityAccess
	default:
		return err
	}
}
