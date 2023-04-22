package types

import (
	entity "ca_parks/core/entity/feature"
)

type SearchFeatureResponse struct {
	Results []entity.Feature
}
