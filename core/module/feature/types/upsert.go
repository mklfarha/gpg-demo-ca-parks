package types

import (
	entity "ca_parks/core/entity/feature"

	"github.com/gofrs/uuid"
	"time"
)

type UpsertRequest struct {
	ID        uuid.UUID
	Name      string
	Status    entity.Status
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
}

type UpsertResponse struct {
	ID uuid.UUID
}
