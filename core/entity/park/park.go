package park

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"
)

type Park struct {
	ID               uuid.UUID       `json:"id"`
	Name             string          `json:"name"`
	MainImage        string          `json:"main_image"`
	Phone            string          `json:"phone"`
	Hours            string          `json:"hours"`
	AllowsDogs       bool            `json:"allows_dogs"`
	Links            LinksCollection `json:"links"`
	Status           Status          `json:"status"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	RecreationAreaID uuid.UUID       `json:"recreation_area_id"`
	UserID           uuid.UUID       `json:"user_id"`
}

func (e Park) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}
