package recreation_area

import (
	"context"

	"ca_parks/core/module/recreation_area/types"
)

func (m *module) FetchRecreationAreaByID(
	ctx context.Context,
	req types.FetchRecreationAreaByIDRequest,
	opts ...Option,
) (types.FetchRecreationAreaByIDResponse, error) {

	models, err := m.repository.FetchRecreationAreaByID(
		ctx,
		req.ID.String(),
	)
	if err != nil {
		return types.FetchRecreationAreaByIDResponse{}, err
	}
	return types.FetchRecreationAreaByIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
