package types

import (
	entity "ca_parks/core/entity/park_has_feature"
)

type FetchParkHasFeatureByStatusRequest struct {
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchParkHasFeatureByStatusResponse struct {
	Results []entity.ParkHasFeature
}
