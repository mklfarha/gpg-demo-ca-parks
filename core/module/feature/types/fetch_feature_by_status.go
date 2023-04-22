package types

import (
	entity "ca_parks/core/entity/feature"
)

type FetchFeatureByStatusRequest struct {
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchFeatureByStatusResponse struct {
	Results []entity.Feature
}
