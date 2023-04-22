package types

import (
	entity "ca_parks/core/entity/user"

	"github.com/gofrs/uuid"
	"time"
)

type UpsertRequest struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Status    entity.Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpsertResponse struct {
	ID uuid.UUID
}
