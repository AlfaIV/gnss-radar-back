package authorization

import "github.com/Gokert/gnss-radar/internal/store"

// IService - сервис для работы c авторизацией
//
//go:generate mockgen -source=$GOFILE -destination=../../../mocks/authorization_service/mock.go
type IAuthorizationService interface {
	Signin(username, password string) (string, error)
	Signup(username, password string) (string, error)
	Authcheck(value string) (bool, error)
	Logout(value string) (bool, error)
}

type AuthorizationService struct {
	authorization store.IAuthorizationStore
}

func NewService(authorization store.IAuthorizationStore) *AuthorizationService {
	return &AuthorizationService{authorization: authorization}
}
