package main

import (
	"context"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"

	ch "challenge"
	"challenge/ent"
	"challenge/ent/migrate"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// accept environment variable as port for server
	// if no port is specified, port 8081 is used
	port := os.Getenv("port")
	if port == "" {
		port = "8081"
	}

	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on specified port.
	srv := handler.NewDefaultServer(ch.NewSchema(client))
	http.Handle("/graphql", srv)
	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), BasicAuth(nil, "secureworks", "supersecret")); err != nil {
		log.Fatal("http server terminated", err)
	}
}

// BasicAuth is a handler func used for basic authentication
func BasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		handler(w, r)
	}
}
