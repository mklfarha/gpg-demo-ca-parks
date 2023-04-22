package park

import (
	"ca_parks/core/module/park/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"context"
	"sync"
)

type Module interface {
	FetchParkByID(ctx context.Context, req types.FetchParkByIDRequest, opts ...Option) (types.FetchParkByIDResponse, error)
	FetchParkByStatus(ctx context.Context, req types.FetchParkByStatusRequest, opts ...Option) (types.FetchParkByStatusResponse, error)
	FetchParkByRecreationAreaID(ctx context.Context, req types.FetchParkByRecreationAreaIDRequest, opts ...Option) (types.FetchParkByRecreationAreaIDResponse, error)
	FetchParkByRecreationAreaIDAndStatus(ctx context.Context, req types.FetchParkByRecreationAreaIDAndStatusRequest, opts ...Option) (types.FetchParkByRecreationAreaIDAndStatusResponse, error)
	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchParkResponse, error)
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
