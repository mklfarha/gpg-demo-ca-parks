package types

import (
	entity "ca_parks/core/entity/user"
)

type FetchUserByEmailRequest struct {
	Email string

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchUserByEmailResponse struct {
	Results []entity.User
}
