package user

import (
	"context"
	"errors"

	"ca_parks/core/module/user/types"
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
		_, err := m.repository.InsertUser(
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

	existing, err := m.repository.FetchUserByID(ctx, req.ID.String())
	if err != nil {
		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		return types.UpsertResponse{}, errors.New("entity not found")
	}
	err = m.repository.UpdateUser(
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

func mapUpsertRequestToInsertParams(req types.UpsertRequest) ca_parksdb.InsertUserParams {
	return ca_parksdb.InsertUserParams{

		ID: custom.GenerateUUID().String(),

		Name: req.Name,

		Email: req.Email,

		Password: custom.Encrypt(req.Password),

		Status: req.Status.ToInt32(),

		CreatedAt: custom.Now(),

		UpdatedAt: custom.Now(),
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest, existing ca_parksdb.User, partial bool) ca_parksdb.UpdateUserParams {
	if !partial {
		return ca_parksdb.UpdateUserParams{

			ID: req.ID.String(),

			Name: req.Name,

			Email: req.Email,

			Password: req.Password,

			Status: req.Status.ToInt32(),

			CreatedAt: existing.CreatedAt,

			UpdatedAt: custom.Now(),
		}
	}

	res := ca_parksdb.UpdateUserParams{}
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

	if req.Email == emptyReq.Email {

		res.Email = existing.Email
	} else {

		res.Email = req.Email

	}

	if req.Password == emptyReq.Password {

		res.Password = existing.Password
	} else {

		res.Password = req.Password

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

	return res

}
