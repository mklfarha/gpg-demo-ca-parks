package feature

import (
	"context"

	"ca_parks/core/module/feature/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchFeatureByStatus(
	ctx context.Context,
	req types.FetchFeatureByStatusRequest,
	opts ...Option,
) (types.FetchFeatureByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchFeatureByStatus(
			ctx,
			ca_parksdb.FetchFeatureByStatusParams{
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchFeatureByStatusResponse{}, err
		}
		return types.FetchFeatureByStatusResponse{
			Results: mapModelsToEntities(models),
		}, nil
	}

	timeFields := []string{
		"created_at", "updated_at",
	}

	orderByFieldFound := false
	for _, tf := range timeFields {
		if tf == req.OrderBy {
			orderByFieldFound = true
			break
		}
	}

	if !orderByFieldFound {
		return types.FetchFeatureByStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchFeatureByStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchFeatureByStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchFeatureByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchFeatureByStatusResponse{}, err
			}
			return types.FetchFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchFeatureByStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchFeatureByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchFeatureByStatusResponse{}, err
			}
			return types.FetchFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchFeatureByStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchFeatureByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchFeatureByStatusResponse{}, err
			}
			return types.FetchFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchFeatureByStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchFeatureByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchFeatureByStatusResponse{}, err
			}
			return types.FetchFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchFeatureByStatusResponse{}, errors.New("could not process request")

}
