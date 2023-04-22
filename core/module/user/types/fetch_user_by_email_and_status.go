package types

import (
	entity "ca_parks/core/entity/user"
)

type FetchUserByEmailAndStatusRequest struct {
	Email  string
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchUserByEmailAndStatusResponse struct {
	Results []entity.User
}
