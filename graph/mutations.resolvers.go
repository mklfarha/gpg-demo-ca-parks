package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"ca_parks/graph/generated"
	"ca_parks/graph/model"

	park "ca_parks/core/module/park/types"

	recreation_area "ca_parks/core/module/recreation_area/types"

	user "ca_parks/core/module/user/types"

	feature "ca_parks/core/module/feature/types"

	park_has_feature "ca_parks/core/module/park_has_feature/types"

	event "ca_parks/core/module/event/types"

	"ca_parks/graph/mapper"
)

func (r *mutationResolver) UpsertPark(
	ctx context.Context,
	input model.ParkInput,
) (*model.Park, error) {
	res, err := r.Core.Park().Upsert(
		ctx,
		mapper.MapParkUpsert(input),
		false,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.Park().
		FetchParkByID(
			ctx,
			park.FetchParkByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapPark(read.Results)[0], err

}

func (r *mutationResolver) UpsertPartialPark(
	ctx context.Context,
	input model.ParkPartialInput,
) (*model.Park, error) {
	res, err := r.Core.Park().Upsert(
		ctx,
		mapper.MapParkUpsertPartial(input),
		true,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.Park().
		FetchParkByID(
			ctx,
			park.FetchParkByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapPark(read.Results)[0], err

}

func (r *mutationResolver) UpsertRecreationArea(
	ctx context.Context,
	input model.RecreationAreaInput,
) (*model.RecreationArea, error) {
	res, err := r.Core.RecreationArea().Upsert(
		ctx,
		mapper.MapRecreationAreaUpsert(input),
		false,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.RecreationArea().
		FetchRecreationAreaByID(
			ctx,
			recreation_area.FetchRecreationAreaByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapRecreationArea(read.Results)[0], err

}

func (r *mutationResolver) UpsertPartialRecreationArea(
	ctx context.Context,
	input model.RecreationAreaPartialInput,
) (*model.RecreationArea, error) {
	res, err := r.Core.RecreationArea().Upsert(
		ctx,
		mapper.MapRecreationAreaUpsertPartial(input),
		true,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.RecreationArea().
		FetchRecreationAreaByID(
			ctx,
			recreation_area.FetchRecreationAreaByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapRecreationArea(read.Results)[0], err

}

func (r *mutationResolver) UpsertUser(
	ctx context.Context,
	input model.UserInput,
) (*model.User, error) {
	res, err := r.Core.User().Upsert(
		ctx,
		mapper.MapUserUpsert(input),
		false,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.User().
		FetchUserByID(
			ctx,
			user.FetchUserByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapUser(read.Results)[0], err

}

func (r *mutationResolver) UpsertPartialUser(
	ctx context.Context,
	input model.UserPartialInput,
) (*model.User, error) {
	res, err := r.Core.User().Upsert(
		ctx,
		mapper.MapUserUpsertPartial(input),
		true,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.User().
		FetchUserByID(
			ctx,
			user.FetchUserByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapUser(read.Results)[0], err

}

func (r *mutationResolver) UpsertFeature(
	ctx context.Context,
	input model.FeatureInput,
) (*model.Feature, error) {
	res, err := r.Core.Feature().Upsert(
		ctx,
		mapper.MapFeatureUpsert(input),
		false,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.Feature().
		FetchFeatureByID(
			ctx,
			feature.FetchFeatureByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapFeature(read.Results)[0], err

}

func (r *mutationResolver) UpsertPartialFeature(
	ctx context.Context,
	input model.FeaturePartialInput,
) (*model.Feature, error) {
	res, err := r.Core.Feature().Upsert(
		ctx,
		mapper.MapFeatureUpsertPartial(input),
		true,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.Feature().
		FetchFeatureByID(
			ctx,
			feature.FetchFeatureByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapFeature(read.Results)[0], err

}

func (r *mutationResolver) UpsertParkHasFeature(
	ctx context.Context,
	input model.ParkHasFeatureInput,
) (*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().Upsert(
		ctx,
		mapper.MapParkHasFeatureUpsert(input),
		false,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.ParkHasFeature().
		FetchParkHasFeatureByID(
			ctx,
			park_has_feature.FetchParkHasFeatureByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapParkHasFeature(read.Results)[0], err

}

func (r *mutationResolver) UpsertPartialParkHasFeature(
	ctx context.Context,
	input model.ParkHasFeaturePartialInput,
) (*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().Upsert(
		ctx,
		mapper.MapParkHasFeatureUpsertPartial(input),
		true,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.ParkHasFeature().
		FetchParkHasFeatureByID(
			ctx,
			park_has_feature.FetchParkHasFeatureByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapParkHasFeature(read.Results)[0], err

}

func (r *mutationResolver) UpsertEvent(
	ctx context.Context,
	input model.EventInput,
) (*model.Event, error) {
	res, err := r.Core.Event().Upsert(
		ctx,
		mapper.MapEventUpsert(input),
		false,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.Event().
		FetchEventByID(
			ctx,
			event.FetchEventByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapEvent(read.Results)[0], err

}

func (r *mutationResolver) UpsertPartialEvent(
	ctx context.Context,
	input model.EventPartialInput,
) (*model.Event, error) {
	res, err := r.Core.Event().Upsert(
		ctx,
		mapper.MapEventUpsertPartial(input),
		true,
	)
	if err != nil {
		return nil, err
	}
	read, err := r.Core.Event().
		FetchEventByID(
			ctx,
			event.FetchEventByIDRequest{
				ID: res.ID,
			},
		)
	if len(read.Results) == 0 {
		return nil, errors.New("entity not found after upsert")
	}
	return mapper.MapEvent(read.Results)[0], err

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
