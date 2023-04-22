package types

import (
	entity "ca_parks/core/entity/park"

	"github.com/gofrs/uuid"
)

type FetchParkByRecreationAreaIDAndStatusRequest struct {
	RecreationAreaID uuid.UUID
	Status           entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkByRecreationAreaIDAndStatusResponse struct {
	Results []entity.Park
}
