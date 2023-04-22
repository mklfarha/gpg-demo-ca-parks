package feature

import (
	"ca_parks/core/module/feature/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"context"
	"sync"
)

type Module interface {
	FetchFeatureByID(ctx context.Context, req types.FetchFeatureByIDRequest, opts ...Option) (types.FetchFeatureByIDResponse, error)
	FetchFeatureByStatus(ctx context.Context, req types.FetchFeatureByStatusRequest, opts ...Option) (types.FetchFeatureByStatusResponse, error)
	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchFeatureResponse, error)
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
