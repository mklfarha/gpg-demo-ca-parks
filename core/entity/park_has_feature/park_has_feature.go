package park_has_feature

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"
)

type ParkHasFeature struct {
	ID        uuid.UUID `json:"id"`
	Details   string    `json:"details"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ParkID    uuid.UUID `json:"park_id"`
	FeatureID uuid.UUID `json:"feature_id"`
}

func (e ParkHasFeature) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}
