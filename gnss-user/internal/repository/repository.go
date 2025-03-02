package statistics_repository

import (
	"context"
	user_domain "gnss-radar/gnss-user/internal"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func (ur *UserRepo) GetUserInfo(ctx context.Context, request user_domain.UserInfoRequest) (user_domain.UserInfoResponse, error) {
    userQuery := `
        SELECT 
            password, 
            login, 
            email, 
            first_name, 
            second_name, 
            role, 
            organization_name 
        FROM user 
        WHERE login = $1;
    `
    var hashedPassword string
    var UserInfo user_domain.UserInfoResponse

    err := ur.pool.QueryRow(ctx, userQuery, request.Login).Scan(
        &hashedPassword,
        &UserInfo.Login,
        &UserInfo.Email,
        &UserInfo.Name,
        &UserInfo.Surname,
        &UserInfo.Role,
        &UserInfo.OrganizationName,
    )
    if err != nil {
        return UserInfo, errors.Wrapf(err, "failed to get user info for user %s", request.Login)
    }

    if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password)); err != nil {
        return UserInfo, errors.Wrapf(err, "failed to validate password for user %s", request.Login)
    }

    apiQuery := `
        SELECT COALESCE(array_agg(api), '{}'::text[]) 
        FROM role_api 
        WHERE role = $1;
    `
    var apis []string
    if err := ur.pool.QueryRow(ctx, apiQuery, UserInfo.Role).Scan(&apis); err != nil {
        return UserInfo, errors.Wrapf(err, "failed to get APIs for role %s", UserInfo.Role)
    }
    
    UserInfo.Api = apis

    return UserInfo, nil
}