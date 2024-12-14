package store

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	sq "github.com/Masterminds/squirrel"
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
	Login            string
	Password         []byte
	Email            string
	OrganizationName string
	FirstName        string
	SecondName       string
}

type SignupResponse struct {
	ID string
}

func (a *AuthorizationStore) Signup(ctx context.Context, params SignupParams) (*model.User, error) {
	query := a.storage.Builder().
		Insert(profileTable).
		SetMap(map[string]any{
			"login":             params.Login,
			"password":          params.Password,
			"email":             params.Email,
			"organization_name": params.OrganizationName,
			"first_name":        params.FirstName,
			"second_name":       params.SecondName,
		}).
		Suffix("RETURNING" + AllProfileTable)

	var response model.User
	if err := a.storage.db.Getx(ctx, &response, query); err != nil {
		return nil, postgresError(err)
	}

	return &response, nil
}

type UserFilter struct {
	IDs    []string
	Logins []string
	Roles  []string
}

func (a *AuthorizationStore) ListUsers(ctx context.Context, filter UserFilter) ([]*model.User, error) {
	query := a.storage.Builder().
		Select(AllProfileTable).
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
	Login    string
	Password []byte
}

func (a *AuthorizationStore) Signin(ctx context.Context, params SigninParams) (*model.User, error) {
	query := a.storage.Builder().
		Select(AllProfileTable).
		From(profileTable).
		Where(map[string]any{
			"login":    params.Login,
			"password": params.Password,
		})

	var user model.User
	if err := a.storage.db.Getx(ctx, &user, query); err != nil {
		return nil, postgresError(err)
	}

	return &user, nil
}

func NewAuthorizationStore(storage *Storage) *AuthorizationStore {
	return &AuthorizationStore{
		storage: storage,
	}
}
