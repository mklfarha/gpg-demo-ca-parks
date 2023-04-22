package types

import (
	entity "ca_parks/core/entity/event"

	"github.com/gofrs/uuid"
)

type FetchEventByParkIDRequest struct {
	ParkID uuid.UUID

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchEventByParkIDResponse struct {
	Results []entity.Event
}
