package park_has_feature

import (
	"context"

	"ca_parks/core/module/park_has_feature/types"
)

func (m *module) FetchParkHasFeatureByID(
	ctx context.Context,
	req types.FetchParkHasFeatureByIDRequest,
	opts ...Option,
) (types.FetchParkHasFeatureByIDResponse, error) {

	models, err := m.repository.FetchParkHasFeatureByID(
		ctx,
		req.ID.String(),
	)
	if err != nil {
		return types.FetchParkHasFeatureByIDResponse{}, err
	}
	return types.FetchParkHasFeatureByIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
