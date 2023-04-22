package types

import (
	entity "ca_parks/core/entity/recreation_area"

	"github.com/gofrs/uuid"
)

type FetchRecreationAreaByIDRequest struct {
	ID uuid.UUID
}

type FetchRecreationAreaByIDResponse struct {
	Results []entity.RecreationArea
}
