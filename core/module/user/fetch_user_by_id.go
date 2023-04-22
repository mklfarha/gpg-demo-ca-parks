package user

import (
	"context"

	"ca_parks/core/module/user/types"
)

func (m *module) FetchUserByID(
	ctx context.Context,
	req types.FetchUserByIDRequest,
	opts ...Option,
) (types.FetchUserByIDResponse, error) {

	models, err := m.repository.FetchUserByID(
		ctx,
		req.ID.String(),
	)
	if err != nil {
		return types.FetchUserByIDResponse{}, err
	}
	return types.FetchUserByIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
