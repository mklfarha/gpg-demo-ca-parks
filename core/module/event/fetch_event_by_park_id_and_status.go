package event

import (
	"context"

	"ca_parks/core/module/event/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchEventByParkIDAndStatus(
	ctx context.Context,
	req types.FetchEventByParkIDAndStatusRequest,
	opts ...Option,
) (types.FetchEventByParkIDAndStatusResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchEventByParkIDAndStatus(
			ctx,
			ca_parksdb.FetchEventByParkIDAndStatusParams{
				ParkID: req.ParkID.String(),
				Status: req.Status.ToInt32(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchEventByParkIDAndStatusResponse{}, err
		}
		return types.FetchEventByParkIDAndStatusResponse{
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
		return types.FetchEventByParkIDAndStatusResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchEventByParkIDAndStatusResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "start_date":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByStartDateASC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByStartDateASCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByStartDateDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByStartDateDESCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "end_date":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByEndDateASC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByEndDateASCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByEndDateDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByEndDateDESCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByCreatedAtASCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByCreatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByUpdatedAtASCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDAndStatusOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDAndStatusOrderedByUpdatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Status: req.Status.ToInt32(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDAndStatusResponse{}, err
			}
			return types.FetchEventByParkIDAndStatusResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchEventByParkIDAndStatusResponse{}, errors.New("could not process request")

}
