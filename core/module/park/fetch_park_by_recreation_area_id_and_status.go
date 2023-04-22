package park

import (
	"context"

	"ca_parks/core/module/park/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkByRecreationAreaIDAndStatus(
	ctx context.Context,
	req types.FetchParkByRecreationAreaIDAndStatusRequest,
	opts ...Option,
) (types.FetchParkByRecreationAreaIDAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkByRecreationAreaIDAndStatus(
			ctx,
			ca_parksdb.FetchParkByRecreationAreaIDAndStatusParams{
				RecreationAreaID: req.RecreationAreaID.String(),
				Status:           req.Status.ToInt32(),
				Offset:           req.Offset,
				Limit:            req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkByRecreationAreaIDAndStatusResponse{}, err
		}
		return types.FetchParkByRecreationAreaIDAndStatusResponse{
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
		return types.FetchParkByRecreationAreaIDAndStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkByRecreationAreaIDAndStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtASCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Status:           req.Status.ToInt32(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDAndStatusResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtDESCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Status:           req.Status.ToInt32(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDAndStatusResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtASCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Status:           req.Status.ToInt32(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDAndStatusResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtDESCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Status:           req.Status.ToInt32(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDAndStatusResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkByRecreationAreaIDAndStatusResponse{}, errors.New("could not process request")

}
