package types

import (
	entity "ca_parks/core/entity/recreation_area"
)

type FetchRecreationAreaByStatusRequest struct {
	Status entity.Status

	Offset  int32
	Limit   int32
	OrderBy string
	Sort    string
}

type FetchRecreationAreaByStatusResponse struct {
	Results []entity.RecreationArea
}
