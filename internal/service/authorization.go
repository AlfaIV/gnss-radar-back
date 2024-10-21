package service

import (
	"context"
	"fmt"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"github.com/Gokert/gnss-radar/internal/store"
	"time"
)

// IAuthorizationService - сервис для работы c авторизацией
//
//go:generate mockgen -source=$GOFILE -destination=../../../mocks/authorization_service/mock.go
type IAuthorizationService interface {
	Signin(ctx context.Context, req SigninRequest) (*model.Session, error)
	Signup(ctx context.Context, req SignupRequest) (*model.User, error)
	ListUsers(ctx context.Context, filter ListUsersFilter) ([]*model.User, error)
	Authcheck(ctx context.Context, value string) (bool, error)
	Logout(ctx context.Context, value string) (bool, error)
}

type AuthorizationService struct {
	authorization store.IAuthorizationStore
	session       store.ISessionStore
}

func NewAuthorizationService(authorization store.IAuthorizationStore, session store.ISessionStore) *AuthorizationService {
	return &AuthorizationService{
		authorization: authorization,
		session:       session,
	}
}

type SigninRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a *AuthorizationService) Signin(ctx context.Context, req SigninRequest) (*model.Session, error) {
	users, err := a.authorization.ListUsers(ctx, store.UserFilter{Logins: []string{req.Login}})
	if err != nil {
		return nil, fmt.Errorf("authorization.ListUsers: %w", err)
	}

	if len(users) == 0 {
		return nil, store.ErrNotFound
	}

	newSession := model.Session{
		Login:     req.Login,
		SID:       utils.RandStringRunes(32),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	sessionAdded, err := a.session.AddSession(ctx, newSession)
	if err != nil {
		return nil, fmt.Errorf("session.AddSession: %w", err)
	}

	if !sessionAdded {
		return nil, store.ErrEntityAlreadyExist
	}

	return &newSession, nil
}

func (a *AuthorizationService) Authcheck(ctx context.Context, value string) (bool, error) {
	result, err := a.session.CheckActiveSession(ctx, value)
	if err != nil {
		return false, fmt.Errorf("session.CheckActiveSession: %w", err)
	}

	return result, nil
}

func (a *AuthorizationService) Logout(ctx context.Context, value string) (bool, error) {
	result, err := a.session.DeleteSession(ctx, value)
	if err != nil {
		return false, fmt.Errorf("session.DeleteSession: %w", err)
	}

	return result, err
}

type SignupRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a *AuthorizationService) Signup(ctx context.Context, req SignupRequest) (*model.User, error) {
	result, err := a.authorization.Signup(ctx, store.SignupParams{Login: req.Login, Password: []byte(req.Password)})
	if err != nil {
		return nil, fmt.Errorf("authorization.Signup: %w", err)
	}

	return result, nil
}

type ListUsersFilter struct {
	IDs    []string `json:"id"`
	Logins []string `json:"login"`
	Role   []string `json:"role"`
}

func (a *AuthorizationService) ListUsers(ctx context.Context, filter ListUsersFilter) ([]*model.User, error) {
	result, err := a.authorization.ListUsers(ctx, store.UserFilter{
		IDs:    filter.IDs,
		Logins: filter.Logins,
		Roles:  filter.Role,
	})
	if err != nil {
		return nil, fmt.Errorf("authorization.ListUsers: %w", err)
	}

	return result, nil
}
