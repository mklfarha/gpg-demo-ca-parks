package park

import (
	"context"

	"ca_parks/core/module/park/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkByStatus(
	ctx context.Context,
	req types.FetchParkByStatusRequest,
	opts ...Option,
) (types.FetchParkByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkByStatus(
			ctx,
			ca_parksdb.FetchParkByStatusParams{
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkByStatusResponse{}, err
		}
		return types.FetchParkByStatusResponse{
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
		return types.FetchParkByStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkByStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkByStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByStatusResponse{}, err
			}
			return types.FetchParkByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkByStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByStatusResponse{}, err
			}
			return types.FetchParkByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkByStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByStatusResponse{}, err
			}
			return types.FetchParkByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkByStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByStatusResponse{}, err
			}
			return types.FetchParkByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkByStatusResponse{}, errors.New("could not process request")

}
