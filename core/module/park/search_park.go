package park

import (
	"context"
	"fmt"

	"ca_parks/core/module/park/types"
	ca_parksdb "ca_parks/core/repository/gen"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchParkResponse, error) {
	models, err := m.repository.SearchPark(
		ctx,
		ca_parksdb.SearchParkParams{
			Name:   fmt.Sprintf("%%%s%%", query),
			Offset: offset,
			Limit:  limit,
		},
	)
	if err != nil {
		return types.SearchParkResponse{}, err
	}
	return types.SearchParkResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
