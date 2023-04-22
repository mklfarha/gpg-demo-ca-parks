package park_has_feature

import (
	"context"

	"ca_parks/core/module/park_has_feature/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkHasFeatureByFeatureID(
	ctx context.Context,
	req types.FetchParkHasFeatureByFeatureIDRequest,
	opts ...Option,
) (types.FetchParkHasFeatureByFeatureIDResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkHasFeatureByFeatureID(
			ctx,
			ca_parksdb.FetchParkHasFeatureByFeatureIDParams{
				FeatureID: req.FeatureID.String(),
				Offset:    req.Offset,
				Limit:     req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkHasFeatureByFeatureIDResponse{}, err
		}
		return types.FetchParkHasFeatureByFeatureIDResponse{
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
		return types.FetchParkHasFeatureByFeatureIDResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkHasFeatureByFeatureIDResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDOrderedByCreatedAtASCParams{
					FeatureID: req.FeatureID.String(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDOrderedByCreatedAtDESCParams{
					FeatureID: req.FeatureID.String(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtASCParams{
					FeatureID: req.FeatureID.String(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtDESCParams{
					FeatureID: req.FeatureID.String(),
					Offset:    req.Offset,
					Limit:     req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByFeatureIDResponse{}, err
			}
			return types.FetchParkHasFeatureByFeatureIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkHasFeatureByFeatureIDResponse{}, errors.New("could not process request")

}
