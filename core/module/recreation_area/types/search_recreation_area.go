package types

import (
	entity "ca_parks/core/entity/recreation_area"
)

type SearchRecreationAreaResponse struct {
	Results []entity.RecreationArea
}
