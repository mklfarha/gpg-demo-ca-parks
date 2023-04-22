package types

import (
	entity "ca_parks/core/entity/event"

	"github.com/gofrs/uuid"
	"time"
)

type UpsertRequest struct {
	ID          uuid.UUID
	Name        string
	Description string
	MainImage   string
	StartDate   time.Time
	EndDate     time.Time
	Status      entity.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ParkID      uuid.UUID
	UserID      uuid.UUID
}

type UpsertResponse struct {
	ID uuid.UUID
}
