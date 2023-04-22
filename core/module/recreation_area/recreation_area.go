package recreation_area

import (
	"ca_parks/core/module/recreation_area/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"context"
	"sync"
)

type Module interface {
	FetchRecreationAreaByID(ctx context.Context, req types.FetchRecreationAreaByIDRequest, opts ...Option) (types.FetchRecreationAreaByIDResponse, error)
	FetchRecreationAreaByStatus(ctx context.Context, req types.FetchRecreationAreaByStatusRequest, opts ...Option) (types.FetchRecreationAreaByStatusResponse, error)
	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchRecreationAreaResponse, error)
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
