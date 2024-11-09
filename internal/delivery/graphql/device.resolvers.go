package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"strconv"
)

// Coords is the resolver for the Coords field.
func (r *deviceResolver) Coords(ctx context.Context, obj *model.Device) (*model.CoordsResults, error) {
	if obj == nil {
		return nil, nil
	}

	return &model.CoordsResults{
		X: strconv.FormatFloat(obj.X, 'f', -1, 64),
		Y: strconv.FormatFloat(obj.Y, 'f', -1, 64),
		Z: strconv.FormatFloat(obj.Z, 'f', -1, 64),
	}, nil
}

// Device returns generated.DeviceResolver implementation.
func (r *Resolver) Device() generated.DeviceResolver { return &deviceResolver{r} }

type deviceResolver struct{ *Resolver }
