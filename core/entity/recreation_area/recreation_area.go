package recreation_area

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"
)

type RecreationArea struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func (e RecreationArea) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}
