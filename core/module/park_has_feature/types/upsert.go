package types

import (
	entity "ca_parks/core/entity/park_has_feature"

	"github.com/gofrs/uuid"
	"time"
)

type UpsertRequest struct {
	ID        uuid.UUID
	Details   string
	Status    entity.Status
	CreatedAt time.Time
	UpdatedAt time.Time
	ParkID    uuid.UUID
	FeatureID uuid.UUID
}

type UpsertResponse struct {
	ID uuid.UUID
}
