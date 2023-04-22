package types

import (
	entity "ca_parks/core/entity/event"

	"github.com/gofrs/uuid"
)

type FetchEventByParkIDAndStatusRequest struct {
	ParkID uuid.UUID
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchEventByParkIDAndStatusResponse struct {
	Results []entity.Event
}
