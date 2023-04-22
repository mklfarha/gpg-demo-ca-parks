package feature

import (
	"context"
	"fmt"

	"ca_parks/core/module/feature/types"
	ca_parksdb "ca_parks/core/repository/gen"
)

func (m *module) Search(
	ctx context.Context,
	query string,
	limit int32,
	offset int32,
) (types.SearchFeatureResponse, error) {
	models, err := m.repository.SearchFeature(
		ctx,
		ca_parksdb.SearchFeatureParams{
			Name:   fmt.Sprintf("%%%s%%", query),
			Offset: offset,
			Limit:  limit,
		},
	)
	if err != nil {
		return types.SearchFeatureResponse{}, err
	}
	return types.SearchFeatureResponse{
		Results: mapModelsToEntities(models),
	}, nil
}
