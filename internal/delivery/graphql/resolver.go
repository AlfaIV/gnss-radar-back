package graphql

import authorization "github.com/Gokert/gnss-radar/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authService authorization.IAuthorizationService
}

func NewResolver(authService authorization.IAuthorizationService) *Resolver {
	return &Resolver{
		authService: authService,
	}
}
