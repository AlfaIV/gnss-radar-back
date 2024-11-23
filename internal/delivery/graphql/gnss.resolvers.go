package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	"github.com/Gokert/gnss-radar/internal/service"
	"github.com/Gokert/gnss-radar/internal/store"
)

// UpdateDevice is the resolver for the updateDevice field.
func (r *gnssMutationsResolver) UpdateDevice(ctx context.Context, obj *model.GnssMutations, input model.UpdateDeviceInput) (*model.UpdateDeviceOutput, error) {
	device, err := r.gnssSevice.UpdateDevice(ctx, service.UpdateDeviceParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		X:           input.Coords.X,
		Y:           input.Coords.Y,
		Z:           input.Coords.Z,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.UpdateDevice %w", err)
	}

	return &model.UpdateDeviceOutput{
		Device: device,
	}, nil
}

// CreateDevice is the resolver for the createDevice field.
func (r *gnssMutationsResolver) CreateDevice(ctx context.Context, obj *model.GnssMutations, input model.CreateDeviceInput) (*model.CreateDeviceOutput, error) {
	device, err := r.gnssSevice.CreateDevice(ctx, service.CreateDeviceParams{
		Name:        input.Name,
		Description: input.Description,
		X:           input.Coords.X,
		Y:           input.Coords.Y,
		Z:           input.Coords.Z,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.CreateDevice %w", err)
	}

	return &model.CreateDeviceOutput{
		Device: device,
	}, nil
}

// DeleteDevice is the resolver for the deleteDevice field.
func (r *gnssMutationsResolver) DeleteDevice(ctx context.Context, obj *model.GnssMutations, input model.DeleteDeviceInput) (*model.DeleteDeviceOutput, error) {
	err := r.gnssSevice.DeleteDevice(ctx, store.DeleteDeviceFilter{Id: input.ID})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.DeleteDevice %w", err)
	}

	return &model.DeleteDeviceOutput{}, nil
}

// CreateTask is the resolver for the createTask field.
func (r *gnssMutationsResolver) CreateTask(ctx context.Context, obj *model.GnssMutations, input model.CreateTaskInput) (*model.CreateTaskOutput, error) {
	task, err := r.gnssSevice.CreateTask(ctx, store.CreateTaskParams{
		SatelliteID:  input.SatelliteID,
		Title:        input.Title,
		Description:  input.Description,
		SignalType:   input.SignalType,
		GroupingType: input.GroupingType,
		StartAt:      input.StartAt,
		EndAt:        input.EndAt,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.CreateTask %w", err)
	}

	return &model.CreateTaskOutput{Task: task}, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *gnssMutationsResolver) UpdateTask(ctx context.Context, obj *model.GnssMutations, input model.UpdateTaskInput) (*model.UpdateTaskOutput, error) {
	task, err := r.gnssSevice.UpdateTask(ctx, store.UpdateTaskParams{
		Id:           input.ID,
		Title:        input.Title,
		Description:  input.Description,
		SatelliteID:  input.SatelliteID,
		SignalType:   input.SignalType,
		GroupingType: input.GroupingType,
		StartAt:      input.StartAt,
		EndAt:        input.EndAt,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.UpdateTask %w", err)
	}

	return &model.UpdateTaskOutput{
		Task: task,
	}, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *gnssMutationsResolver) DeleteTask(ctx context.Context, obj *model.GnssMutations, input model.DeleteTaskInput) (*model.DeleteTaskOutput, error) {
	err := r.gnssSevice.DeleteTask(ctx, store.DeleteTaskFilter{Id: input.ID})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.DeleteTask %w", err)
	}

	return &model.DeleteTaskOutput{}, nil
}

// CreateSatellite is the resolver for the createSatellite field.
func (r *gnssMutationsResolver) CreateSatellite(ctx context.Context, obj *model.GnssMutations, input model.CreateSatelliteInput) (*model.CreateSatelliteOutput, error) {
	satellite, err := r.gnssSevice.CreateSatellite(ctx, store.CreateSatelliteParams{
		ExternalSatelliteId: input.ExternalSatelliteID,
		SatelliteName:       input.SatelliteName,
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.CreateSatellite %w", err)
	}

	return &model.CreateSatelliteOutput{Satellite: satellite}, nil
}

// Gnss is the resolver for the gnss field.
func (r *mutationResolver) Gnss(ctx context.Context) (*model.GnssMutations, error) {
	return &model.GnssMutations{}, nil
}

// ListGnss is the resolver for the listGnss field.
func (r *queryResolver) ListGnss(ctx context.Context, filter model.GNSSFilter, page int, perPage int) (*model.GNSSPagination, error) {
	if filter.Coordinates == nil {
		return nil, nil
	}

	gnssList, err := r.gnssSevice.ListGnss(ctx, service.ListGnssRequest{
		X: filter.Coordinates.X,
		Y: filter.Coordinates.Y,
		Z: filter.Coordinates.Z,
		Paginator: model.Paginator{
			Page:    uint64(page),
			PerPage: uint64(perPage),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.ListGnss: %w", err)
	}

	return &model.GNSSPagination{
		Items: utils.SerializerGnssCoords(gnssList),
	}, nil
}

// ListDevice is the resolver for the listDevice field.
func (r *queryResolver) ListDevice(ctx context.Context, filter model.DeviceFilter, page int, perPage int) (*model.DevicePagination, error) {
	devices, err := r.gnssSevice.ListDevice(ctx, service.ListDeviceFilter{
		Ids:    filter.Ids,
		Names:  filter.Names,
		Tokens: filter.Tokens,
		Paginator: model.Paginator{
			Page:    uint64(page),
			PerPage: uint64(perPage),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.ListDevice: %w", err)
	}

	return &model.DevicePagination{Items: devices}, nil
}

// ListTask is the resolver for the listTask field.
func (r *queryResolver) ListTask(ctx context.Context, filter model.TaskFilter, page int, perPage int) (*model.TaskPagination, error) {
	tasks, err := r.gnssSevice.ListTasks(ctx, store.ListTasksFilter{
		Ids:          filter.Ids,
		SatelliteIds: filter.SatelliteIds,
		SignalType:   filter.SignalType,
		GroupingType: filter.GroupingType,
		StartAt:      filter.StartAt,
		EndAt:        filter.EndAt,
		Paginator: model.Paginator{
			Page:    uint64(page),
			PerPage: uint64(perPage),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.ListTasks: %w", err)
	}

	return &model.TaskPagination{Items: tasks}, nil
}

// Rinexlist is the resolver for the Rinexlist field.
func (r *queryResolver) Rinexlist(ctx context.Context, input *model.RinexInput, page int, perPage int) (*model.RinexPagination, error) {
	gnssRinex, err := r.gnssSevice.RinexList(ctx, service.RinexRequest{})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.Rinexlist: %w", err)
	}
	return &model.RinexPagination{
		Items: gnssRinex,
	}, nil
}

// ListSatellites is the resolver for the listSatellites field.
func (r *queryResolver) ListSatellites(ctx context.Context, filter model.SatellitesFilter, page int, perPage int) (*model.SatellitesPagination, error) {
	satellites, err := r.gnssSevice.ListSatellites(ctx, store.ListSatellitesFilter{
		Ids:                  filter.IDS,
		ExternalSatelliteIds: filter.ExternalSatelliteIds,
		SatelliteName:        filter.SatelliteNames,
		Paginator:            model.Paginator{Page: uint64(page), PerPage: uint64(perPage)},
	})
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.ListSatellites: %w", err)
	}

	return &model.SatellitesPagination{Items: satellites}, nil
}

// ListMeasurements is the resolver for the listMeasurements field.
func (r *queryResolver) ListMeasurements(ctx context.Context, filter model.MeasurementsFilter, page int, perPage int) (*model.MeasurementsPagination, error) {
	measurements, err := r.gnssSevice.ListMeasurements(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("gnssSevice.ListMeasurements: %w", err)
	}

	return &model.MeasurementsPagination{Items: measurements}, nil
}

// GnssMutations returns generated.GnssMutationsResolver implementation.
func (r *Resolver) GnssMutations() generated.GnssMutationsResolver { return &gnssMutationsResolver{r} }

type gnssMutationsResolver struct{ *Resolver }
