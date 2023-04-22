package types

import (
	entity "ca_parks/core/entity/event"

	"github.com/gofrs/uuid"
)

type FetchEventByIDRequest struct {
	ID uuid.UUID
}

type FetchEventByIDResponse struct {
	Results []entity.Event
}
