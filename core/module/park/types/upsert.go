package types

import (
	entity "ca_parks/core/entity/park"

	"github.com/gofrs/uuid"
	"time"
)

type UpsertRequest struct {
	ID               uuid.UUID
	Name             string
	MainImage        string
	Phone            string
	Hours            string
	AllowsDogs       bool
	Links            entity.LinksCollection
	Status           entity.Status
	CreatedAt        time.Time
	UpdatedAt        time.Time
	RecreationAreaID uuid.UUID
	UserID           uuid.UUID
}

type UpsertResponse struct {
	ID uuid.UUID
}
