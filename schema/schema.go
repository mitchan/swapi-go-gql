package schema

import (
	"github.com/graphql-go/graphql"
)

type Character struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
	Gender string `json:"gender"`
}

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

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/*
		   curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
		*/
		"people": &graphql.Field{
			Type:        graphql.NewList(characterType),
			Description: "List of characters",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return characters, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var SwapiSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
