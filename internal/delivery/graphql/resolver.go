package graphql

import "github.com/Gokert/gnss-radar/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authService service.IAuthorizationService
	gnssSevice  service.IGnss
}

func NewResolver(service2 *service.Service) *Resolver {
	return &Resolver{
		authService: service2.GetAuthorizationService(),
		gnssSevice:  service2.GetGnssService(),
	}
}
