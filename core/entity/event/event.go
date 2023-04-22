package event

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MainImage   string    `json:"main_image"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ParkID      uuid.UUID `json:"park_id"`
	UserID      uuid.UUID `json:"user_id"`
}

func (e Event) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}
