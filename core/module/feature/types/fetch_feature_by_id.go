package types

import (
	entity "ca_parks/core/entity/feature"

	"github.com/gofrs/uuid"
)

type FetchFeatureByIDRequest struct {
	ID uuid.UUID
}

type FetchFeatureByIDResponse struct {
	Results []entity.Feature
}
