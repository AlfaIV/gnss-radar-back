package service

import "github.com/Gokert/gnss-radar/internal/store"

type Service struct {
	authorization *AuthorizationService
	gnss          *GnssService
}

func NewService(auth store.IAuthorizationStore, session store.ISessionStore, gnss store.IGnssStore) *Service {
	return &Service{
		authorization: NewAuthorizationService(auth, session),
		gnss:          NewGnssService(gnss),
	}
}

func (s *Service) GetAuthorizationService() *AuthorizationService {
	return s.authorization
}

func (s *Service) GetGnssService() *GnssService {
	return s.gnss
}
