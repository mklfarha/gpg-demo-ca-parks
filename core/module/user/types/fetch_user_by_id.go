package types

import (
	entity "ca_parks/core/entity/user"

	"github.com/gofrs/uuid"
)

type FetchUserByIDRequest struct {
	ID uuid.UUID
}

type FetchUserByIDResponse struct {
	Results []entity.User
}
