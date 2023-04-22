package recreation_area

import (
	"context"
	"fmt"

	"ca_parks/core/module/recreation_area/types"
	ca_parksdb "ca_parks/core/repository/gen"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchRecreationAreaResponse, error) {
	models, err := m.repository.SearchRecreationArea(
		ctx,
		ca_parksdb.SearchRecreationAreaParams{
			Name:   fmt.Sprintf("%%%s%%", query),
			Offset: offset,
			Limit:  limit,
		},
	)
	if err != nil {
		return types.SearchRecreationAreaResponse{}, err
	}
	return types.SearchRecreationAreaResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
