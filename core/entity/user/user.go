package user

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e User) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}
