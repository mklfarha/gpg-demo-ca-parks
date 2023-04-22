package types

import (
	entity "ca_parks/core/entity/park_has_feature"

	"github.com/gofrs/uuid"
)

type FetchParkHasFeatureByFeatureIDAndStatusRequest struct {
	FeatureID uuid.UUID
	Status    entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkHasFeatureByFeatureIDAndStatusResponse struct {
	Results []entity.ParkHasFeature
}
