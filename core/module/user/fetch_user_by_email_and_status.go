package user

import (
	"context"

	"ca_parks/core/module/user/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchUserByEmailAndStatus(
	ctx context.Context,
	req types.FetchUserByEmailAndStatusRequest,
	opts ...Option,
) (types.FetchUserByEmailAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchUserByEmailAndStatus(
			ctx,
			ca_parksdb.FetchUserByEmailAndStatusParams{
				Email:  req.Email,
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchUserByEmailAndStatusResponse{}, err
		}
		return types.FetchUserByEmailAndStatusResponse{
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
		return types.FetchUserByEmailAndStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchUserByEmailAndStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchUserByEmailAndStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchUserByEmailAndStatusOrderedByCreatedAtASCParams{
					Email:  req.Email,
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchUserByEmailAndStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchUserByEmailAndStatusOrderedByCreatedAtDESCParams{
					Email:  req.Email,
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchUserByEmailAndStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchUserByEmailAndStatusOrderedByUpdatedAtASCParams{
					Email:  req.Email,
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchUserByEmailAndStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchUserByEmailAndStatusOrderedByUpdatedAtDESCParams{
					Email:  req.Email,
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailAndStatusResponse{}, err
			}
			return types.FetchUserByEmailAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchUserByEmailAndStatusResponse{}, errors.New("could not process request")

}
