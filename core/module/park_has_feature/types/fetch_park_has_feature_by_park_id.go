package types

import (
	entity "ca_parks/core/entity/park_has_feature"

	"github.com/gofrs/uuid"
)

type FetchParkHasFeatureByParkIDRequest struct {
	ParkID uuid.UUID

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkHasFeatureByParkIDResponse struct {
	Results []entity.ParkHasFeature
}
