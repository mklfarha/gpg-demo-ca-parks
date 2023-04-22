package feature

import (
	"context"

	"ca_parks/core/module/feature/types"
)

func (m *module) FetchFeatureByID(
	ctx context.Context,
	req types.FetchFeatureByIDRequest,
	opts ...Option,
) (types.FetchFeatureByIDResponse, error) {

	models, err := m.repository.FetchFeatureByID(
		ctx,
		req.ID.String(),
	)
	if err != nil {
		return types.FetchFeatureByIDResponse{}, err
	}
	return types.FetchFeatureByIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
