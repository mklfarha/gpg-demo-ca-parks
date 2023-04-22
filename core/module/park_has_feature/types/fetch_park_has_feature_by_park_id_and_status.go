package types

import (
	entity "ca_parks/core/entity/park_has_feature"

	"github.com/gofrs/uuid"
)

type FetchParkHasFeatureByParkIDAndStatusRequest struct {
	ParkID uuid.UUID
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkHasFeatureByParkIDAndStatusResponse struct {
	Results []entity.ParkHasFeature
}
