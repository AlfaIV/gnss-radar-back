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

func (ur *UserRepo) GetUserInfoById(ctx context.Context, userId string) (user_domain.UserInfoResponse, error) {
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
        WHERE id = $1;
    `
	var hashedPassword string
	var UserInfo user_domain.UserInfoResponse

	err := ur.pool.QueryRow(ctx, userQuery, userId).Scan(
		&hashedPassword,
		&UserInfo.Login,
		&UserInfo.Email,
		&UserInfo.Name,
		&UserInfo.Surname,
		&UserInfo.Role,
		&UserInfo.OrganizationName,
	)
	if err != nil {
		return UserInfo, errors.Wrapf(err, "failed to get user info for user with id %s", userId)
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

func (ur *UserRepo) CreateUser(ctx context.Context, request user_domain.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	if err != nil {
		return errors.Wrapf(err, "failed to generate hashed password for %s", request.Login)
	}

	createUserQuery := "insert into user (login, email, password, first_name, second_name, organization_name, role) values ($1, $2, $3, $4, $5, $6, $7)"

	if _, err := ur.pool.Query(ctx, createUserQuery, request.Login, request.Email, hashedPassword, request.Name, request.Surname, request.OrganizationName, request.Role); err != nil {
		return errors.Wrapf(err, "failed to create account for %s", request.Login)
	}

	return nil
}

func (ur *UserRepo) ValidatePermissions(ctx context.Context, userId string, api string) (bool, error) {
	validatePermissionsQuery := `
        SELECT EXISTS(
            SELECT 1
            FROM user u
            INNER JOIN role_api ra ON u.role = ra.role
            WHERE u.login = $1 
            AND ra.api = $2
        );
    `

	var exists bool
	err := ur.pool.QueryRow(
		ctx,
		validatePermissionsQuery,
		userId,
		api,
	).Scan(&exists)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, errors.Wrapf(err, "failed to validate permissions for user %s", userId)
	}

	return exists, nil
}

func (ur *UserRepo) ResolveUserSignUp(ctx context.Context, userLogin string, resolution string) error {
	resolutionQuery := `
	UPDATE user SET status = 1$ WHERE login = 2$;
    `

	if _, err := ur.pool.Query(ctx, resolutionQuery, resolution, userLogin); err != nil {
		return errors.Wrapf(err, "failed to make resolution for %s", userLogin)
	}

	return nil
}

func (ur *UserRepo) ChangeUserPermissions(ctx context.Context, userLogin string, userRole string) error {
	resolutionQuery := `
	UPDATE user SET role = 1$ WHERE login = 2$;
    `
	if _, err := ur.pool.Query(ctx, resolutionQuery, userRole, userLogin); err != nil {
		return errors.Wrapf(err, "failed to change permissions for %s", userLogin)
	}

	return nil
}

func (ur *UserRepo) GetSignUpRequestions(ctx context.Context, params user_domain.PaginatedRequest) ([]user_domain.UserSignUpRequestion, error) {
	query := `
        SELECT 
            login, 
            email, 
            first_name, 
            second_name,
			organization_name
        FROM user
        WHERE status = 'PENDING'
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2;
    `

	offset := (params.Page - 1) * params.Size

	rows, err := ur.pool.Query(
		ctx,
		query,
		params.Size,
		offset,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get requestions")
	}
	defer rows.Close()

	var users []user_domain.UserSignUpRequestion

	for rows.Next() {
		var user user_domain.UserSignUpRequestion
		err := rows.Scan(
			&user.Login,
			&user.Email,
			&user.Name,
			&user.Surname,
			&user.OrganizationName,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error during rows iteration")
	}

	return users, nil
}

func (ur *UserRepo) GetUserForAdmin(ctx context.Context, params user_domain.PaginatedRequest) ([]user_domain.UserForAdmin, error) {
	query := `
        SELECT 
            login, 
            email, 
            first_name, 
            second_name,
			organization_name,
			role
        FROM user
        WHERE status <> 'PENDING'
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2;
    `

	offset := (params.Page - 1) * params.Size

	rows, err := ur.pool.Query(
		ctx,
		query,
		params.Size,
		offset,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get requestions")
	}
	defer rows.Close()

	var users []user_domain.UserForAdmin

	for rows.Next() {
		var user user_domain.UserForAdmin
		err := rows.Scan(
			&user.Login,
			&user.Email,
			&user.Name,
			&user.Surname,
			&user.OrganizationName,
			&user.Role,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error during rows iteration")
	}

	return users, nil
}
