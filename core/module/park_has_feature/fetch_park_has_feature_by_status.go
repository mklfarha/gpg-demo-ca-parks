package park_has_feature

import (
	"context"

	"ca_parks/core/module/park_has_feature/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkHasFeatureByStatus(
	ctx context.Context,
	req types.FetchParkHasFeatureByStatusRequest,
	opts ...Option,
) (types.FetchParkHasFeatureByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkHasFeatureByStatus(
			ctx,
			ca_parksdb.FetchParkHasFeatureByStatusParams{
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkHasFeatureByStatusResponse{}, err
		}
		return types.FetchParkHasFeatureByStatusResponse{
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
		return types.FetchParkHasFeatureByStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkHasFeatureByStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkHasFeatureByStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkHasFeatureByStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkHasFeatureByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkHasFeatureByStatusResponse{}, err
			}
			return types.FetchParkHasFeatureByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkHasFeatureByStatusResponse{}, errors.New("could not process request")

}
