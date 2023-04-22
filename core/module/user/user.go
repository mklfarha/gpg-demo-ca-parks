package user

import (
	"ca_parks/core/module/user/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"context"
	"sync"
)

type Module interface {
	FetchUserByID(ctx context.Context, req types.FetchUserByIDRequest, opts ...Option) (types.FetchUserByIDResponse, error)
	FetchUserByEmail(ctx context.Context, req types.FetchUserByEmailRequest, opts ...Option) (types.FetchUserByEmailResponse, error)
	FetchUserByStatus(ctx context.Context, req types.FetchUserByStatusRequest, opts ...Option) (types.FetchUserByStatusResponse, error)
	FetchUserByEmailAndStatus(ctx context.Context, req types.FetchUserByEmailAndStatusRequest, opts ...Option) (types.FetchUserByEmailAndStatusResponse, error)
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
