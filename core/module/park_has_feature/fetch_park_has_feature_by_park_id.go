package park_has_feature

import (
	"context"

	"ca_parks/core/module/park_has_feature/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkHasFeatureByParkID(
	ctx context.Context,
	req types.FetchParkHasFeatureByParkIDRequest,
	opts ...Option,
) (types.FetchParkHasFeatureByParkIDResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkHasFeatureByParkID(
			ctx,
			ca_parksdb.FetchParkHasFeatureByParkIDParams{
				ParkID: req.ParkID.String(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkHasFeatureByParkIDResponse{}, err
		}
		return types.FetchParkHasFeatureByParkIDResponse{
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
		return types.FetchParkHasFeatureByParkIDResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkHasFeatureByParkIDResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByParkIDOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDOrderedByCreatedAtASCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByParkIDOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDOrderedByCreatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByParkIDOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDOrderedByUpdatedAtASCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByParkIDOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDOrderedByUpdatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkHasFeatureByParkIDResponse{}, errors.New("could not process request")

}
