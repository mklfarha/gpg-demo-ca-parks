package event

import (
	"context"
	"errors"

	"ca_parks/core/module/event/types"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"

	"ca_parks/custom"
)

func (m *module) Upsert(
	ctx context.Context,
	req types.UpsertRequest,
	partial bool,
	opts ...Option,
) (types.UpsertResponse, error) {
	if req.ID == uuid.Nil {
		params := mapUpsertRequestToInsertParams(req)
		_, err := m.repository.InsertEvent(
			ctx,
			params,
		)
		if err != nil {
			return types.UpsertResponse{}, err
		}

		return types.UpsertResponse{
			ID: uuid.FromStringOrNil(params.ID),
		}, nil
	}

	existing, err := m.repository.FetchEventByID(ctx, req.ID.String())
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		return types.UpsertResponse{}, errors.New("entity not found")
	}
	err = m.repository.UpdateEvent(
		ctx,
		mapUpsertRequestToUpdateParams(req, existing[0], partial),
	)
	if err != nil {
		return types.UpsertResponse{}, err
	}

	return types.UpsertResponse{
		ID: req.ID,
	}, nil
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) ca_parksdb.InsertEventParams {
	return ca_parksdb.InsertEventParams{

		ID: custom.GenerateUUID().String(),

		Name: req.Name,

		Description: req.Description,

		MainImage: req.MainImage,

		StartDate: req.StartDate,

		EndDate: req.EndDate,

		Status: req.Status.ToInt32(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		ParkID: req.ParkID.String(),

		UserID: req.UserID.String(),
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing ca_parksdb.Event, partial bool) ca_parksdb.UpdateEventParams {
	if !partial {
		return ca_parksdb.UpdateEventParams{

			ID: req.ID.String(),

			Name: req.Name,

			Description: req.Description,

			MainImage: req.MainImage,

			StartDate: req.StartDate,

			EndDate: req.EndDate,

			Status: req.Status.ToInt32(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			ParkID: req.ParkID.String(),

			UserID: req.UserID.String(),
		}
	}

	res := ca_parksdb.UpdateEventParams{}
	emptyReq := types.UpsertRequest{}

	if req.ID == emptyReq.ID {

		res.ID = existing.ID
	} else {

		res.ID = req.ID.String()

	}

	if req.Name == emptyReq.Name {

		res.Name = existing.Name
	} else {

		res.Name = req.Name

	}

	if req.Description == emptyReq.Description {

		res.Description = existing.Description
	} else {

		res.Description = req.Description

	}

	if req.MainImage == emptyReq.MainImage {

		res.MainImage = existing.MainImage
	} else {

		res.MainImage = req.MainImage

	}

	if req.StartDate == emptyReq.StartDate {

		res.StartDate = existing.StartDate
	} else {

		res.StartDate = req.StartDate

	}

	if req.EndDate == emptyReq.EndDate {

		res.EndDate = existing.EndDate
	} else {

		res.EndDate = req.EndDate

	}

	if req.Status == emptyReq.Status {

		res.Status = existing.Status
	} else {

		res.Status = req.Status.ToInt32()

	}

	if req.CreatedAt == emptyReq.CreatedAt {

		res.CreatedAt = existing.CreatedAt
	} else {

		res.CreatedAt = existing.CreatedAt

	}

	if req.UpdatedAt == emptyReq.UpdatedAt {

		res.UpdatedAt = existing.UpdatedAt
	} else {

		res.UpdatedAt = custom.Now()

	}

	if req.ParkID == emptyReq.ParkID {

		res.ParkID = existing.ParkID
	} else {

		res.ParkID = req.ParkID.String()

	}

	if req.UserID == emptyReq.UserID {

		res.UserID = existing.UserID
	} else {

		res.UserID = req.UserID.String()

	}

	return res

}
