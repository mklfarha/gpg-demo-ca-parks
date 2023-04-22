package auth

import (
	"context"
	"net/http"
	"time"

	"ca_parks/core"
)

var jwtKey = []byte("some_jwt_key")

type TokenResponse struct {
	Token   string
	Expires time.Time
}

type Interface interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	ClaimsFromHeader(w http.ResponseWriter, r *http.Request) (*Claims, error)
}

type Implementation struct {
	core *core.Implementation
}

func New(ctx context.Context, c *core.Implementation) (Interface, error) {
	return &Implementation{
		core: c,
	}, nil

}
