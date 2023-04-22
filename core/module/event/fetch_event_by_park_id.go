package event

import (
	"context"

	"ca_parks/core/module/event/types"

	ca_parksdb "ca_parks/core/repository/gen"
	"errors"
)

func (m *module) FetchEventByParkID(
	ctx context.Context,
	req types.FetchEventByParkIDRequest,
	opts ...Option,
) (types.FetchEventByParkIDResponse, error) {

	if req.OrderBy == "" {
		models, err := m.repository.FetchEventByParkID(
			ctx,
			ca_parksdb.FetchEventByParkIDParams{
				ParkID: req.ParkID.String(),
				Offset: req.Offset,
				Limit:  req.Limit,
			},
		)
		if err != nil {
			return types.FetchEventByParkIDResponse{}, err
		}
		return types.FetchEventByParkIDResponse{
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
		return types.FetchEventByParkIDResponse{}, errors.New("order by field not found")
	}

	sort := "ASC"
	if req.Sort != "" {
		if req.Sort == "ASC" || req.Sort == "DESC" {
			sort = req.Sort
		} else {
			return types.FetchEventByParkIDResponse{}, errors.New("invalid sort value only ASC or DESC are valid")
		}
	}

	switch req.OrderBy {

	case "start_date":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDOrderedByStartDateASC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByStartDateASCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDOrderedByStartDateDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByStartDateDESCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "end_date":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDOrderedByEndDateASC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByEndDateASCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDOrderedByEndDateDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByEndDateDESCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "created_at":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDOrderedByCreatedAtASC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByCreatedAtASCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDOrderedByCreatedAtDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByCreatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	case "updated_at":
		if sort == "ASC" {
			models, err := m.repository.FetchEventByParkIDOrderedByUpdatedAtASC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByUpdatedAtASCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		} else {
			models, err := m.repository.FetchEventByParkIDOrderedByUpdatedAtDESC(
				ctx,
				ca_parksdb.FetchEventByParkIDOrderedByUpdatedAtDESCParams{
					ParkID: req.ParkID.String(),
					Offset: req.Offset,
					Limit:  req.Limit,
				},
			)
			if err != nil {
				return types.FetchEventByParkIDResponse{}, err
			}
			return types.FetchEventByParkIDResponse{
				Results: mapModelsToEntities(models),
			}, nil
		}

	}

	return types.FetchEventByParkIDResponse{}, errors.New("could not process request")

}
