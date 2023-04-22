package types

import (
	entity "ca_parks/core/entity/park"

	"github.com/gofrs/uuid"
)

type FetchParkByIDRequest struct {
	ID uuid.UUID
}

type FetchParkByIDResponse struct {
	Results []entity.Park
}
