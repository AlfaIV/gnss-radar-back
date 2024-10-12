package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
)

// Authorization is the resolver for the authorization field.
func (r *mutationResolver) Authorization(ctx context.Context) (*model.AuthorizationMutations, error) {
	return &model.AuthorizationMutations{}, nil
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context, input *model.TestInput) (*model.TestOutput, error) {
	panic(fmt.Errorf("not implemented: Test - test"))
}
