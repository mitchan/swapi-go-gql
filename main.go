package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	swapi "github.com/mitchan/swapi-go-gql/schema"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        swapi.SwapiSchema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	h := handler.New(&handler.Config{
		Schema:     &swapi.SwapiSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
