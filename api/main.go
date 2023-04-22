package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app, err := New()
	if err != nil {
		panic(err)
	}

	http.Handle("/gplay", app.basicAuth(playground.Handler("GraphQL playground", "/query")))
	http.HandleFunc("/signin", app.auth.SignIn)
	http.HandleFunc("/validate", app.auth.Validate)
	http.HandleFunc("/refresh", app.auth.Refresh)
	http.Handle("/query", app.serverHandlerFunc())
	http.Handle("/upload", app.serverHandlerUploadFunc())

	log.Printf(`
		Port: %s
			/gplay
			/signin
			/refresh
			/query
			/upload
			`, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
