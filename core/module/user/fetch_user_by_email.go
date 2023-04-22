package user

import (
	"context"

	"ca_parks/core/module/user/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchUserByEmail(
	ctx context.Context,
	req types.FetchUserByEmailRequest,
	opts ...Option,
) (types.FetchUserByEmailResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchUserByEmail(
			ctx,
			ca_parksdb.FetchUserByEmailParams{
				Email:  req.Email,
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchUserByEmailResponse{}, err
		}
		return types.FetchUserByEmailResponse{
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
		return types.FetchUserByEmailResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchUserByEmailResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchUserByEmailOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchUserByEmailOrderedByCreatedAtASCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchUserByEmailOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchUserByEmailOrderedByCreatedAtDESCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchUserByEmailOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchUserByEmailOrderedByUpdatedAtASCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchUserByEmailOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchUserByEmailOrderedByUpdatedAtDESCParams{
					Email:  req.Email,
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchUserByEmailResponse{}, err
			}
			return types.FetchUserByEmailResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchUserByEmailResponse{}, errors.New("could not process request")

}
