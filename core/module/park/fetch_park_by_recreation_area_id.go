package park

import (
	"context"

	"ca_parks/core/module/park/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchParkByRecreationAreaID(
	ctx context.Context,
	req types.FetchParkByRecreationAreaIDRequest,
	opts ...Option,
) (types.FetchParkByRecreationAreaIDResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchParkByRecreationAreaID(
			ctx,
			ca_parksdb.FetchParkByRecreationAreaIDParams{
				RecreationAreaID: req.RecreationAreaID.String(),
				Offset:           req.Offset,
				Limit:            req.Limit,
			},
		)
		if err != nil {
			return types.FetchParkByRecreationAreaIDResponse{}, err
		}
		return types.FetchParkByRecreationAreaIDResponse{
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
		return types.FetchParkByRecreationAreaIDResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchParkByRecreationAreaIDResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkByRecreationAreaIDOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDOrderedByCreatedAtASCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkByRecreationAreaIDOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDOrderedByCreatedAtDESCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchParkByRecreationAreaIDOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDOrderedByUpdatedAtASCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchParkByRecreationAreaIDOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchParkByRecreationAreaIDOrderedByUpdatedAtDESCParams{
					RecreationAreaID: req.RecreationAreaID.String(),
					Offset:           req.Offset,
					Limit:            req.Limit,
				},
			)
			if err != nil {
				return types.FetchParkByRecreationAreaIDResponse{}, err
			}
			return types.FetchParkByRecreationAreaIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchParkByRecreationAreaIDResponse{}, errors.New("could not process request")

}
