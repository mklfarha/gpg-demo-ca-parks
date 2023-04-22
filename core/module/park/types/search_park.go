package types

import (
	entity "ca_parks/core/entity/park"
)

type SearchParkResponse struct {
	Results []entity.Park
}
