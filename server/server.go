package main

import (
	_ "github.com/99designs/gqlgen"
)

func main() {
	// router := http.NewServeMux()

	// router.Handle("/", playground.Handler("Dataloader", "/query"))
	// router.Handle("/query", handler.NewDefaultServer(
	// 	dataloader.NewExecutableSchema(dataloader.Config{Resolvers: &dataloader.Resolver{}}),
	// ))

	// log.Println("connect to http://localhost:8082/ for graphql playground")
	// log.Fatal(http.ListenAndServe(":8082", dataloader.LoaderMiddleware(router)))
}