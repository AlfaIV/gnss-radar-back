package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"fmt"

	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"github.com/Gokert/gnss-radar/internal/service"
	"github.com/Gokert/gnss-radar/internal/store"
)

// Signup is the resolver for the signup field.
func (r *authorizationMutationsResolver) Signup(ctx context.Context, obj *model.AuthorizationMutations, input model.SignupInput) (*model.SignupOutput, error) {
	if err := utils.ValidateSignup(&input); err != nil {
		return nil, fmt.Errorf("validation.ValidateStruct %w", err)
	}

	user, err := r.authService.Signup(ctx, service.SignupRequest{
		Login:            input.Login,
		Password:         input.Password,
		Email:            input.Email,
		OrganizationName: input.OrganizationName,
		FirstName:        input.FirstName,
		SecondName:       input.SecondName,
	})
	if err != nil {
		switch {
		case errors.Is(err, store.ErrEntityAlreadyExist):
			return nil, model.ErrorAlreadyExists
		default:
			return nil, fmt.Errorf("authService.Signup %w", err)
		}
	}

	return &model.SignupOutput{UserInfo: user}, nil
}

// Signin is the resolver for the signin field.
func (r *authorizationMutationsResolver) Signin(ctx context.Context, obj *model.AuthorizationMutations, input model.SigninInput) (*model.SigninOutput, error) {
	if err := utils.ValidateSignin(&input); err != nil {
		return nil, fmt.Errorf("validation.ValidateStruct %w", err)
	}

	session, user, err := r.authService.Signin(ctx, service.SigninRequest{Login: input.Login, Password: input.Password})
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return nil, model.ErrorUserNotFound
		default:
			return nil, fmt.Errorf("authService.Signin %w", err)
		}
	}

	utils.SetCookie(ctx, session.SID)

	return &model.SigninOutput{UserInfo: user}, nil
}

// Logout is the resolver for the logout field.
func (r *authorizationMutationsResolver) Logout(ctx context.Context, obj *model.AuthorizationMutations, input *model.LogoutInput) (*model.LogoutOutput, error) {
	cookie, err := utils.GetRequest(ctx).Cookie("session_id")
	if err != nil {
		return nil, model.ErrorCookieNotFound
	}

	result, err := r.authService.Logout(ctx, cookie.Value)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrorNotAuthorized):
			return nil, model.ErrorUserNotFound
		default:
			return nil, fmt.Errorf("authService.Logout: %w", err)
		}
	}

	if !result {
		return nil, model.ErrorInternalError
	}

	utils.RemoveCookie(ctx)

	return &model.LogoutOutput{}, nil
}

// Authorization is the resolver for the authorization field.
func (r *mutationResolver) Authorization(ctx context.Context) (*model.AuthorizationMutations, error) {
	return &model.AuthorizationMutations{}, nil
}

// Authcheck is the resolver for the authcheck field.
func (r *queryResolver) Authcheck(ctx context.Context, input *model.AuthcheckInput) (*model.AuthcheckOutput, error) {
	cookie, err := utils.GetRequest(ctx).Cookie("session_id")
	if err != nil {
		return nil, model.ErrorCookieNotFound
	}

	result, user, err := r.authService.Authcheck(ctx, cookie.Value)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return nil, model.ErrorUserNotFound
		default:
			return nil, fmt.Errorf("authService.Authcheck: %w", err)
		}
	}

	if !result {
		return nil, model.ErrorNotAuthorized
	}

	return &model.AuthcheckOutput{UserInfo: user}, nil
}

// AuthorizationMutations returns generated.AuthorizationMutationsResolver implementation.
func (r *Resolver) AuthorizationMutations() generated.AuthorizationMutationsResolver {
	return &authorizationMutationsResolver{r}
}

type authorizationMutationsResolver struct{ *Resolver }
