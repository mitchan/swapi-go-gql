package schema

import (
	"encoding/json"

	"github.com/mitchan/swapi-go-gql/types"
	"github.com/mitchan/swapi-go-gql/utils"
)

var Schema = `
schema {
	query: Query
}

type Query {
	people: [Character!]
	planets: [Planet!]
	films: [Film!]
}

type Character {
	name: String!
	height: String!
	mass: String!
	gender: String!
	homeworld: Planet
	films: [Film!]
}

type Planet {
	name: String!
	climate: String!
	terrain: String!
	population: String!
	residents: [Character!]
	films: [Film!]
}

type Film {
	title: String!
	openingCrawl: String!
	director: String!
	releaseDate: String!
	characters: [Character!]
	planets: [Planet!]
}
`

var endpoints = []string{"people", "planets", "films"}

func PrefetchData() {
	for _, endpoint := range endpoints {
		switch endpoint {
		case "people":
			go loadEndpoint(endpoint, &types.People)

		case "planets":
			go loadEndpoint(endpoint, &types.Planets)

		case "films":
			go loadEndpoint(endpoint, &types.Films)
		}
	}
}

func loadEndpoint(endpoint string, i interface{}) error {
	bytes, err := utils.LoadData(endpoint)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return err
	}

	return nil
}
