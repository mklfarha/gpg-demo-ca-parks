package types

import (
	entity "ca_parks/core/entity/user"
)

type FetchUserByStatusRequest struct {
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchUserByStatusResponse struct {
	Results []entity.User
}
