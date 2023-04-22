package park_has_feature

import (
	"context"

	"ca_parks/core/module/park_has_feature/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkHasFeatureByParkIDAndStatus(
	ctx context.Context,
	req types.FetchParkHasFeatureByParkIDAndStatusRequest,
	opts ...Option,
) (types.FetchParkHasFeatureByParkIDAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkHasFeatureByParkIDAndStatus(
			ctx,
			ca_parksdb.FetchParkHasFeatureByParkIDAndStatusParams{
				ParkID: req.ParkID.String(),
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, err
		}
		return types.FetchParkHasFeatureByParkIDAndStatusResponse{
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
		return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtASCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtASCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkHasFeatureByParkIDAndStatusResponse{}, errors.New("could not process request")

}
