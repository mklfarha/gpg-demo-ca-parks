package event

import (
	"ca_parks/core/module/event/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"context"
	"sync"
)

type Module interface {
	FetchEventByID(ctx context.Context, req types.FetchEventByIDRequest, opts ...Option) (types.FetchEventByIDResponse, error)
	FetchEventByStatus(ctx context.Context, req types.FetchEventByStatusRequest, opts ...Option) (types.FetchEventByStatusResponse, error)
	FetchEventByParkID(ctx context.Context, req types.FetchEventByParkIDRequest, opts ...Option) (types.FetchEventByParkIDResponse, error)
	FetchEventByParkIDAndStatus(ctx context.Context, req types.FetchEventByParkIDAndStatusRequest, opts ...Option) (types.FetchEventByParkIDAndStatusResponse, error)
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
