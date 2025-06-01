package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"go-graphql/graph/generated"

	"go-graphql/graph"
)

func main() {
	resolver := graph.NewResolver()

	// Create the GraphQL server with the generated schema + resolvers
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)

	// Serve GraphQL playground at "/"
	http.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))

	// Serve GraphQL API at "/graphql"
	http.Handle("/graphql", srv)

	log.Println("🚀 Server ready at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
