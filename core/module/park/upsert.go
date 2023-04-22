package park

import (
	"context"
	"errors"

	"ca_parks/core/module/park/types"
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
		_, err := m.repository.InsertPark(
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

	existing, err := m.repository.FetchParkByID(ctx, req.ID.String())
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		return types.UpsertResponse{}, errors.New("entity not found")
	}
	err = m.repository.UpdatePark(
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

func mapUpsertRequestToInsertParams(req types.UpsertRequest) ca_parksdb.InsertParkParams {
	return ca_parksdb.InsertParkParams{

		ID: custom.GenerateUUID().String(),

		Name: req.Name,

		MainImage: req.MainImage,

		Phone: req.Phone,

		Hours: req.Hours,

		AllowsDogs: req.AllowsDogs,

		Links: req.Links.LinksToJSON(),

		Status: req.Status.ToInt32(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		RecreationAreaID: req.RecreationAreaID.String(),

		UserID: req.UserID.String(),
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing ca_parksdb.Park, partial bool) ca_parksdb.UpdateParkParams {
	if !partial {
		return ca_parksdb.UpdateParkParams{

			ID: req.ID.String(),

			Name: req.Name,

			MainImage: req.MainImage,

			Phone: req.Phone,

			Hours: req.Hours,

			AllowsDogs: req.AllowsDogs,

			Links: req.Links.LinksToJSON(),

			Status: req.Status.ToInt32(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			RecreationAreaID: req.RecreationAreaID.String(),

			UserID: req.UserID.String(),
		}
	}

	res := ca_parksdb.UpdateParkParams{}
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

	if req.MainImage == emptyReq.MainImage {

		res.MainImage = existing.MainImage
	} else {

		res.MainImage = req.MainImage

	}

	if req.Phone == emptyReq.Phone {

		res.Phone = existing.Phone
	} else {

		res.Phone = req.Phone

	}

	if req.Hours == emptyReq.Hours {

		res.Hours = existing.Hours
	} else {

		res.Hours = req.Hours

	}

	if req.AllowsDogs == emptyReq.AllowsDogs {

		res.AllowsDogs = existing.AllowsDogs
	} else {

		res.AllowsDogs = req.AllowsDogs

	}

	if len(req.Links) == 0 {

		res.Links = existing.Links
	} else {

		res.Links = req.Links.LinksToJSON()

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

	if req.RecreationAreaID == emptyReq.RecreationAreaID {

		res.RecreationAreaID = existing.RecreationAreaID
	} else {

		res.RecreationAreaID = req.RecreationAreaID.String()

	}

	if req.UserID == emptyReq.UserID {

		res.UserID = existing.UserID
	} else {

		res.UserID = req.UserID.String()

	}

	return res

}
