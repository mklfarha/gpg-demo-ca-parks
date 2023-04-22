package event

import (
	"context"

	"ca_parks/core/module/event/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchEventByStatus(
	ctx context.Context,
	req types.FetchEventByStatusRequest,
	opts ...Option,
) (types.FetchEventByStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchEventByStatus(
			ctx,
			ca_parksdb.FetchEventByStatusParams{
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchEventByStatusResponse{}, err
		}
		return types.FetchEventByStatusResponse{
			Results: mapModelsToEntities(models),
		}, nil
	}

	timeFields := []string{
		"start_date", "end_date", "created_at", "updated_at",
	}

	orderByFieldFound := false
	for _, tf := range timeFields {
		if tf == req.OrderBy {
			orderByFieldFound = true
			break
		}
	}

	if !orderByFieldFound {
		return types.FetchEventByStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchEventByStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "start_date":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByStatusOrderedByStartDateASC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByStartDateASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByStatusOrderedByStartDateDESC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByStartDateDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "end_date":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByStatusOrderedByEndDateASC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByEndDateASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByStatusOrderedByEndDateDESC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByEndDateDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByCreatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByCreatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByUpdatedAtASCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchEventByStatusOrderedByUpdatedAtDESCParams{
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByStatusResponse{}, err
			}
			return types.FetchEventByStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchEventByStatusResponse{}, errors.New("could not process request")

}
