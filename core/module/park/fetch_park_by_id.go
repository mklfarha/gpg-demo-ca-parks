package park

import (
	"context"

	"ca_parks/core/module/park/types"
)

func (m *module) FetchParkByID(
	ctx context.Context,
	req types.FetchParkByIDRequest,
	opts ...Option,
) (types.FetchParkByIDResponse, error) {

	models, err := m.repository.FetchParkByID(
		ctx,
		req.ID.String(),
	)
	if err != nil {
		return types.FetchParkByIDResponse{}, err
	}
	return types.FetchParkByIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
