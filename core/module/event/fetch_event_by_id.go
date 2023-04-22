package event

import (
	"context"

	"ca_parks/core/module/event/types"
)

func (m *module) FetchEventByID(
	ctx context.Context,
	req types.FetchEventByIDRequest,
	opts ...Option,
) (types.FetchEventByIDResponse, error) {

	models, err := m.repository.FetchEventByID(
		ctx,
		req.ID.String(),
	)
	if err != nil {
		return types.FetchEventByIDResponse{}, err
	}
	return types.FetchEventByIDResponse{
		Results: mapModelsToEntities(models),
	}, nil

}
