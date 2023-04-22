package main

import (
	"context"

	"ca_parks/auth"
	"ca_parks/config"
	"ca_parks/core"
	"ca_parks/graph"
	"ca_parks/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
)

type basicAuthCredentials struct {
	username string
	password string
}

type application struct {
	core   *core.Implementation
	auth   auth.Interface
	server *handler.Server
	bac    basicAuthCredentials
}

func New() (application, error) {
	ctx := context.Background()
	config := config.New()
	c, err := core.New(ctx, config)
	if err != nil {
		return application{}, err
	}

	auth, err := auth.New(ctx, c)
	if err != nil {
		return application{}, err
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Core: *c,
				}}))

	return application{
		core:   c,
		auth:   auth,
		server: srv,
		bac: basicAuthCredentials{
			username: "admin",
			password: "4dm1n",
		},
	}, nil
}
