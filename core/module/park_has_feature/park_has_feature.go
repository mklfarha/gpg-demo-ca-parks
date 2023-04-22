package park_has_feature

import (
	"ca_parks/core/module/park_has_feature/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"context"
	"sync"
)

type Module interface {
	FetchParkHasFeatureByID(ctx context.Context, req types.FetchParkHasFeatureByIDRequest, opts ...Option) (types.FetchParkHasFeatureByIDResponse, error)
	FetchParkHasFeatureByStatus(ctx context.Context, req types.FetchParkHasFeatureByStatusRequest, opts ...Option) (types.FetchParkHasFeatureByStatusResponse, error)
	FetchParkHasFeatureByParkID(ctx context.Context, req types.FetchParkHasFeatureByParkIDRequest, opts ...Option) (types.FetchParkHasFeatureByParkIDResponse, error)
	FetchParkHasFeatureByParkIDAndStatus(ctx context.Context, req types.FetchParkHasFeatureByParkIDAndStatusRequest, opts ...Option) (types.FetchParkHasFeatureByParkIDAndStatusResponse, error)
	FetchParkHasFeatureByFeatureID(ctx context.Context, req types.FetchParkHasFeatureByFeatureIDRequest, opts ...Option) (types.FetchParkHasFeatureByFeatureIDResponse, error)
	FetchParkHasFeatureByFeatureIDAndStatus(ctx context.Context, req types.FetchParkHasFeatureByFeatureIDAndStatusRequest, opts ...Option) (types.FetchParkHasFeatureByFeatureIDAndStatusResponse, error)
	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
}

type module struct {
	repository ca_parksdb.Querier
	mu         sync.Mutex
}

func New(repo ca_parksdb.Querier) Module {
	return &module{
		repository: repo,
	}
}
