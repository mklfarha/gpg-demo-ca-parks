package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ca_parks/graph/generated"
	"ca_parks/graph/model"
    
    park "ca_parks/core/module/park/types"
    parkentity "ca_parks/core/entity/park"
    
    recreation_area "ca_parks/core/module/recreation_area/types"
    recreation_areaentity "ca_parks/core/entity/recreation_area"
    
    user "ca_parks/core/module/user/types"
    userentity "ca_parks/core/entity/user"
    
    feature "ca_parks/core/module/feature/types"
    featureentity "ca_parks/core/entity/feature"
    
    park_has_feature "ca_parks/core/module/park_has_feature/types"
    park_has_featureentity "ca_parks/core/entity/park_has_feature"
    
    event "ca_parks/core/module/event/types"
    evententity "ca_parks/core/entity/event"
    	 
    "github.com/gofrs/uuid"  
    "ca_parks/graph/mapper" 	
)


func (r *queryResolver) ParkByID(
    ctx context.Context, 
    id *string,
    
) ([]*model.Park, error) {
	res, err := r.Core.Park().
    FetchParkByID(
        ctx, 
        park.FetchParkByIDRequest{ 
            ID: mapper.UuidFromPointerString(id),
            
	    },
    )
	return mapper.MapPark(res.Results), err
}

func (r *queryResolver) ParkByStatus(
    ctx context.Context, 
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Park, error) {
	res, err := r.Core.Park().
    FetchParkByStatus(
        ctx, 
        park.FetchParkByStatusRequest{ 
            Status: parkentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapPark(res.Results), err
}

func (r *queryResolver) ParkByRecreationAreaID(
    ctx context.Context, 
    recreation_area_id string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Park, error) {
	res, err := r.Core.Park().
    FetchParkByRecreationAreaID(
        ctx, 
        park.FetchParkByRecreationAreaIDRequest{ 
            RecreationAreaID: uuid.FromStringOrNil(recreation_area_id),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapPark(res.Results), err
}

func (r *queryResolver) ParkByRecreationAreaIDAndStatus(
    ctx context.Context, 
    recreation_area_id string,
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Park, error) {
	res, err := r.Core.Park().
    FetchParkByRecreationAreaIDAndStatus(
        ctx, 
        park.FetchParkByRecreationAreaIDAndStatusRequest{ 
            RecreationAreaID: uuid.FromStringOrNil(recreation_area_id),
            Status: parkentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapPark(res.Results), err
}



func (r *queryResolver) SearchPark(
    ctx context.Context,
    query string,
    limit *int, 
    offset *int,
) ([]*model.Park, error) {
    res, err := r.Core.Park().
        Search(
            ctx, 
            query,
            resolveLimit(limit),
            resolveOffset(offset),
        )
	return mapper.MapPark(res.Results), err
}








func (r *queryResolver) RecreationAreaByID(
    ctx context.Context, 
    id *string,
    
) ([]*model.RecreationArea, error) {
	res, err := r.Core.RecreationArea().
    FetchRecreationAreaByID(
        ctx, 
        recreation_area.FetchRecreationAreaByIDRequest{ 
            ID: mapper.UuidFromPointerString(id),
            
	    },
    )
	return mapper.MapRecreationArea(res.Results), err
}

func (r *queryResolver) RecreationAreaByStatus(
    ctx context.Context, 
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.RecreationArea, error) {
	res, err := r.Core.RecreationArea().
    FetchRecreationAreaByStatus(
        ctx, 
        recreation_area.FetchRecreationAreaByStatusRequest{ 
            Status: recreation_areaentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapRecreationArea(res.Results), err
}



func (r *queryResolver) SearchRecreationArea(
    ctx context.Context,
    query string,
    limit *int, 
    offset *int,
) ([]*model.RecreationArea, error) {
    res, err := r.Core.RecreationArea().
        Search(
            ctx, 
            query,
            resolveLimit(limit),
            resolveOffset(offset),
        )
	return mapper.MapRecreationArea(res.Results), err
}








func (r *queryResolver) UserByID(
    ctx context.Context, 
    id *string,
    
) ([]*model.User, error) {
	res, err := r.Core.User().
    FetchUserByID(
        ctx, 
        user.FetchUserByIDRequest{ 
            ID: mapper.UuidFromPointerString(id),
            
	    },
    )
	return mapper.MapUser(res.Results), err
}

func (r *queryResolver) UserByEmail(
    ctx context.Context, 
    email string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.User, error) {
	res, err := r.Core.User().
    FetchUserByEmail(
        ctx, 
        user.FetchUserByEmailRequest{ 
            Email: email,
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapUser(res.Results), err
}

func (r *queryResolver) UserByStatus(
    ctx context.Context, 
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.User, error) {
	res, err := r.Core.User().
    FetchUserByStatus(
        ctx, 
        user.FetchUserByStatusRequest{ 
            Status: userentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapUser(res.Results), err
}

func (r *queryResolver) UserByEmailAndStatus(
    ctx context.Context, 
    email string,
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.User, error) {
	res, err := r.Core.User().
    FetchUserByEmailAndStatus(
        ctx, 
        user.FetchUserByEmailAndStatusRequest{ 
            Email: email,
            Status: userentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapUser(res.Results), err
}










func (r *queryResolver) FeatureByID(
    ctx context.Context, 
    id *string,
    
) ([]*model.Feature, error) {
	res, err := r.Core.Feature().
    FetchFeatureByID(
        ctx, 
        feature.FetchFeatureByIDRequest{ 
            ID: mapper.UuidFromPointerString(id),
            
	    },
    )
	return mapper.MapFeature(res.Results), err
}

func (r *queryResolver) FeatureByStatus(
    ctx context.Context, 
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Feature, error) {
	res, err := r.Core.Feature().
    FetchFeatureByStatus(
        ctx, 
        feature.FetchFeatureByStatusRequest{ 
            Status: featureentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapFeature(res.Results), err
}



func (r *queryResolver) SearchFeature(
    ctx context.Context,
    query string,
    limit *int, 
    offset *int,
) ([]*model.Feature, error) {
    res, err := r.Core.Feature().
        Search(
            ctx, 
            query,
            resolveLimit(limit),
            resolveOffset(offset),
        )
	return mapper.MapFeature(res.Results), err
}








func (r *queryResolver) ParkHasFeatureByID(
    ctx context.Context, 
    id *string,
    
) ([]*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().
    FetchParkHasFeatureByID(
        ctx, 
        park_has_feature.FetchParkHasFeatureByIDRequest{ 
            ID: mapper.UuidFromPointerString(id),
            
	    },
    )
	return mapper.MapParkHasFeature(res.Results), err
}

func (r *queryResolver) ParkHasFeatureByStatus(
    ctx context.Context, 
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().
    FetchParkHasFeatureByStatus(
        ctx, 
        park_has_feature.FetchParkHasFeatureByStatusRequest{ 
            Status: park_has_featureentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapParkHasFeature(res.Results), err
}

func (r *queryResolver) ParkHasFeatureByParkID(
    ctx context.Context, 
    park_id string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().
    FetchParkHasFeatureByParkID(
        ctx, 
        park_has_feature.FetchParkHasFeatureByParkIDRequest{ 
            ParkID: uuid.FromStringOrNil(park_id),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapParkHasFeature(res.Results), err
}

func (r *queryResolver) ParkHasFeatureByParkIDAndStatus(
    ctx context.Context, 
    park_id string,
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().
    FetchParkHasFeatureByParkIDAndStatus(
        ctx, 
        park_has_feature.FetchParkHasFeatureByParkIDAndStatusRequest{ 
            ParkID: uuid.FromStringOrNil(park_id),
            Status: park_has_featureentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapParkHasFeature(res.Results), err
}

func (r *queryResolver) ParkHasFeatureByFeatureID(
    ctx context.Context, 
    feature_id string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().
    FetchParkHasFeatureByFeatureID(
        ctx, 
        park_has_feature.FetchParkHasFeatureByFeatureIDRequest{ 
            FeatureID: uuid.FromStringOrNil(feature_id),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapParkHasFeature(res.Results), err
}

func (r *queryResolver) ParkHasFeatureByFeatureIDAndStatus(
    ctx context.Context, 
    feature_id string,
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.ParkHasFeature, error) {
	res, err := r.Core.ParkHasFeature().
    FetchParkHasFeatureByFeatureIDAndStatus(
        ctx, 
        park_has_feature.FetchParkHasFeatureByFeatureIDAndStatusRequest{ 
            FeatureID: uuid.FromStringOrNil(feature_id),
            Status: park_has_featureentity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapParkHasFeature(res.Results), err
}










func (r *queryResolver) EventByID(
    ctx context.Context, 
    id *string,
    
) ([]*model.Event, error) {
	res, err := r.Core.Event().
    FetchEventByID(
        ctx, 
        event.FetchEventByIDRequest{ 
            ID: mapper.UuidFromPointerString(id),
            
	    },
    )
	return mapper.MapEvent(res.Results), err
}

func (r *queryResolver) EventByStatus(
    ctx context.Context, 
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Event, error) {
	res, err := r.Core.Event().
    FetchEventByStatus(
        ctx, 
        event.FetchEventByStatusRequest{ 
            Status: evententity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapEvent(res.Results), err
}

func (r *queryResolver) EventByParkID(
    ctx context.Context, 
    park_id string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Event, error) {
	res, err := r.Core.Event().
    FetchEventByParkID(
        ctx, 
        event.FetchEventByParkIDRequest{ 
            ParkID: uuid.FromStringOrNil(park_id),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapEvent(res.Results), err
}

func (r *queryResolver) EventByParkIDAndStatus(
    ctx context.Context, 
    park_id string,
    status string,
    
    limit *int,
    offset *int,
    orderBy *string,
    sort *string,
    
) ([]*model.Event, error) {
	res, err := r.Core.Event().
    FetchEventByParkIDAndStatus(
        ctx, 
        event.FetchEventByParkIDAndStatusRequest{ 
            ParkID: uuid.FromStringOrNil(park_id),
            Status: evententity.StatusFromString(status),
            
            Limit: resolveLimit(limit),
            Offset: resolveOffset(offset),
            OrderBy: stringPointerToString(orderBy),
            Sort: stringPointerToString(sort),
            
	    },
    )
	return mapper.MapEvent(res.Results), err
}










// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func resolveLimit(limit *int) int32 {
    if limit == nil {
        return int32(10)
    }
    return int32(*limit)
}

func resolveOffset(offset *int) int32 {
    if offset == nil {
        return int32(0)
    }
    return int32(*offset)
}

func stringPointerToString(in *string) string {
    if in == nil {
        return ""
    }
    return *in
}
