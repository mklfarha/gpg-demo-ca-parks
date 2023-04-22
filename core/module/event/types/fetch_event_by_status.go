package types

import (
	entity "ca_parks/core/entity/event"
)

type FetchEventByStatusRequest struct {
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchEventByStatusResponse struct {
	Results []entity.Event
}
