package types

import (
	entity "ca_parks/core/entity/park_has_feature"

	"github.com/gofrs/uuid"
)

type FetchParkHasFeatureByIDRequest struct {
	ID uuid.UUID
}

type FetchParkHasFeatureByIDResponse struct {
	Results []entity.ParkHasFeature
}
