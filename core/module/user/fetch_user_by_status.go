package user

import (
	"context"

	"ca_parks/core/module/user/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchUserByStatus(
	ctx context.Context,
	req types.FetchUserByStatusRequest,
	opts ...Option,
) (types.FetchUserByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchUserByStatus(
			ctx,
			ca_parksdb.FetchUserByStatusParams{
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchUserByStatusResponse{}, err
		}
		return types.FetchUserByStatusResponse{
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
		return types.FetchUserByStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchUserByStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchUserByStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchUserByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByStatusResponse{}, err
			}
			return types.FetchUserByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchUserByStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchUserByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByStatusResponse{}, err
			}
			return types.FetchUserByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchUserByStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchUserByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByStatusResponse{}, err
			}
			return types.FetchUserByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchUserByStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchUserByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByStatusResponse{}, err
			}
			return types.FetchUserByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchUserByStatusResponse{}, errors.New("could not process request")

}
