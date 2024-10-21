package store

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	profileTable = "profile"
)

type IAuthorizationStore interface {
	Signin(ctx context.Context, params SigninParams) (*model.User, error)
	Signup(ctx context.Context, params SignupParams) (*model.User, error)
	ListUsers(ctx context.Context, filter UserFilter) ([]*model.User, error)
}

type AuthorizationStore struct {
	storage *Storage
}

type SignupParams struct {
	Login    string `json:"login"`
	Password []byte `json:"password"`
}

type SignupResponse struct {
	ID string `json:"id"`
}

func (a *AuthorizationStore) Signup(ctx context.Context, params SignupParams) (*model.User, error) {
	query := a.storage.Builder().
		Insert(profileTable).
		SetMap(map[string]any{
			"login":    params.Login,
			"password": params.Password,
		}).
		Suffix("RETURNING id, login, role")

	var response model.User
	if err := a.storage.db.Getx(ctx, &response, query); err != nil {
		return nil, postgresError(err)
	}

	return &response, nil
}

type UserFilter struct {
	IDs    []string `json:"id"`
	Logins []string `json:"login"`
	Roles  []string `json:"role"`
}

func (a *AuthorizationStore) ListUsers(ctx context.Context, filter UserFilter) ([]*model.User, error) {
	query := a.storage.Builder().
		Select("id, login, role").
		From(profileTable)

	if len(filter.IDs) > 0 {
		query = query.Where(sq.Eq{"id": filter.IDs})
	}
	if len(filter.Logins) > 0 {
		query = query.Where(sq.Eq{"login": filter.Logins})
	}
	if len(filter.Roles) > 0 {
		query = query.Where(sq.Eq{"roles": filter.Roles})
	}

	var users []*model.User
	if err := a.storage.db.Selectx(ctx, &users, query); err != nil {
		return nil, postgresError(err)
	}

	return users, nil
}

type SigninParams struct {
	Login    string `json:"login"`
	Password []byte `json:"password"`
}

func (a *AuthorizationStore) Signin(ctx context.Context, params SigninParams) (*model.User, error) {
	query := a.storage.Builder().
		Select("id, login, role").
		From(profileTable).
		Where(map[string]any{
			"login":    params.Login,
			"password": params.Password,
		})

	var user *model.User
	if err := a.storage.db.Getx(ctx, &user, query); err != nil {
		return nil, postgresError(err)
	}

	return user, nil
}

func NewAuthorizationStore(storage *Storage) *AuthorizationStore {
	return &AuthorizationStore{
		storage: storage,
	}
}
