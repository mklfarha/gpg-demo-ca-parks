package recreation_area

import (
	"context"

	"ca_parks/core/module/recreation_area/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchRecreationAreaByStatus(
	ctx context.Context,
	req types.FetchRecreationAreaByStatusRequest,
	opts ...Option,
) (types.FetchRecreationAreaByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchRecreationAreaByStatus(
			ctx,
			ca_parksdb.FetchRecreationAreaByStatusParams{
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchRecreationAreaByStatusResponse{}, err
		}
		return types.FetchRecreationAreaByStatusResponse{
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
		return types.FetchRecreationAreaByStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchRecreationAreaByStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchRecreationAreaByStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchRecreationAreaByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchRecreationAreaByStatusResponse{}, err
			}
			return types.FetchRecreationAreaByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchRecreationAreaByStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchRecreationAreaByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchRecreationAreaByStatusResponse{}, err
			}
			return types.FetchRecreationAreaByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchRecreationAreaByStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchRecreationAreaByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchRecreationAreaByStatusResponse{}, err
			}
			return types.FetchRecreationAreaByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchRecreationAreaByStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchRecreationAreaByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchRecreationAreaByStatusResponse{}, err
			}
			return types.FetchRecreationAreaByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchRecreationAreaByStatusResponse{}, errors.New("could not process request")

}
