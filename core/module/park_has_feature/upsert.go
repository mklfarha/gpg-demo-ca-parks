package park_has_feature

import (
	"context"
	"errors"

	"ca_parks/core/module/park_has_feature/types"
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
		_, err := m.repository.InsertParkHasFeature(
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

	existing, err := m.repository.FetchParkHasFeatureByID(ctx, req.ID.String())
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		return types.UpsertResponse{}, errors.New("entity not found")
	}
	err = m.repository.UpdateParkHasFeature(
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

func mapUpsertRequestToInsertParams(req types.UpsertRequest) ca_parksdb.InsertParkHasFeatureParams {
	return ca_parksdb.InsertParkHasFeatureParams{

		ID: custom.GenerateUUID().String(),

		Details: req.Details,

		Status: req.Status.ToInt32(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),

		ParkID: req.ParkID.String(),

		FeatureID: req.FeatureID.String(),
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing ca_parksdb.ParkHasFeature, partial bool) ca_parksdb.UpdateParkHasFeatureParams {
	if !partial {
		return ca_parksdb.UpdateParkHasFeatureParams{

			ID: req.ID.String(),

			Details: req.Details,

			Status: req.Status.ToInt32(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),

			ParkID: req.ParkID.String(),

			FeatureID: req.FeatureID.String(),
		}
	}

	res := ca_parksdb.UpdateParkHasFeatureParams{}
	emptyReq := types.UpsertRequest{}

	if req.ID == emptyReq.ID {

		res.ID = existing.ID
	} else {

		res.ID = req.ID.String()

	}

	if req.Details == emptyReq.Details {

		res.Details = existing.Details
	} else {

		res.Details = req.Details

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

	if req.FeatureID == emptyReq.FeatureID {

		res.FeatureID = existing.FeatureID
	} else {

		res.FeatureID = req.FeatureID.String()

	}

	return res

}
