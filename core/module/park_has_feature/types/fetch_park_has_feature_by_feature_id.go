package types

import (
	entity "ca_parks/core/entity/park_has_feature"

	"github.com/gofrs/uuid"
)

type FetchParkHasFeatureByFeatureIDRequest struct {
	FeatureID uuid.UUID

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkHasFeatureByFeatureIDResponse struct {
	Results []entity.ParkHasFeature
}
