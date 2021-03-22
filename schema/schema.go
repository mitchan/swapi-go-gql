package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var characterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "People",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"height": &graphql.Field{
			Type: graphql.String,
		},
		"mass": &graphql.Field{
			Type: graphql.Boolean,
		},
		"gender": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var characters []Character

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"people": &graphql.Field{
			Type:        graphql.NewList(characterType),
			Description: "List of characters",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := http.Get("https://swapi.dev/api/people")
				if err != nil {
					return nil, err
				}

				bytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}

				var people AllPeople

				if err := json.Unmarshal(bytes, &people); err != nil {
					fmt.Println("Error parsing json", err)
					return nil, err
				}

				return people.People, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var SwapiSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
