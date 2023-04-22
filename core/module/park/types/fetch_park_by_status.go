package types

import (
	entity "ca_parks/core/entity/park"
)

type FetchParkByStatusRequest struct {
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkByStatusResponse struct {
	Results []entity.Park
}
