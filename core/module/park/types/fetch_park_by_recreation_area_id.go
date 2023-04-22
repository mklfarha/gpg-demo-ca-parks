package types

import (
	entity "ca_parks/core/entity/park"

	"github.com/gofrs/uuid"
)

type FetchParkByRecreationAreaIDRequest struct {
	RecreationAreaID uuid.UUID

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkByRecreationAreaIDResponse struct {
	Results []entity.Park
}
