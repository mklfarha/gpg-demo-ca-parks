package park_has_feature

import (
	"context"

	"ca_parks/core/module/park_has_feature/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkHasFeatureByFeatureIDAndStatus(
	ctx context.Context,
	req types.FetchParkHasFeatureByFeatureIDAndStatusRequest,
	opts ...Option,
) (types.FetchParkHasFeatureByFeatureIDAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkHasFeatureByFeatureIDAndStatus(
			ctx,
			ca_parksdb.FetchParkHasFeatureByFeatureIDAndStatusParams{
				FeatureID: req.FeatureID.String(),
				Status:    req.Status.ToInt32(),
				Offset:    req.Offset,
				Limit:     req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, err
		}
		return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{
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
		return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtASCParams{
					FeatureID: req.FeatureID.String(),
					Status:    req.Status.ToInt32(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtDESCParams{
					FeatureID: req.FeatureID.String(),
					Status:    req.Status.ToInt32(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtASCParams{
					FeatureID: req.FeatureID.String(),
					Status:    req.Status.ToInt32(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtDESCParams{
					FeatureID: req.FeatureID.String(),
					Status:    req.Status.ToInt32(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkHasFeatureByFeatureIDAndStatusResponse{}, errors.New("could not process request")

}
