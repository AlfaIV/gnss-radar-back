package service

import "github.com/Gokert/gnss-radar/internal/store"

type Service struct {
	authorization *AuthorizationService
}

func NewService(auth store.IAuthorizationStore, session store.ISessionStore) *Service {
	return &Service{authorization: NewAuthorizationService(auth, session)}
}

func (s *Service) GetAuthorizationService() *AuthorizationService {
	return s.authorization
}
