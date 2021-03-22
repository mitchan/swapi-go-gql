package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

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

var planetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Planet",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"climate": &graphql.Field{
			Type: graphql.String,
		},
		"terrain": &graphql.Field{
			Type: graphql.Boolean,
		},
		"residents": &graphql.Field{
			Type: graphql.NewList(characterType),
		},
	},
})

const baseUrl = "https://swapi.dev/api/"

var people AllPeople
var planets AllPlanets

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"people": &graphql.Field{
			Type:        graphql.NewList(characterType),
			Description: "List of characters",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if people.People == nil {
					// load from api
					resp, err := http.Get(baseUrl + "people")
					if err != nil {
						return nil, err
					}

					bytes, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return nil, err
					}

					if err := json.Unmarshal(bytes, &people); err != nil {
						fmt.Println("Error parsing json", err)
						return nil, err
					}
				}

				return people.People, nil
			},
		},
		"planets": &graphql.Field{
			Type:        graphql.NewList(planetType),
			Description: "List of characters",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if planets.Planets == nil {
					// load from api
					resp, err := http.Get(baseUrl + "planets")
					if err != nil {
						return nil, err
					}

					bytes, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return nil, err
					}

					if err := json.Unmarshal(bytes, &planets); err != nil {
						fmt.Println("Error parsing json", err)
						return nil, err
					}
				}

				return planets.Planets, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var SwapiSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
