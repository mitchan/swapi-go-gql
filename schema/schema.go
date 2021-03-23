package schema

import (
	"encoding/json"

	"github.com/mitchan/swapi-go-gql/utils"
)

// type Planet {
// 	name: String
// 	climate: String
// 	terrain: String
// 	population: String
// 	urls: [String]
// }

var Schema = `
schema {
	query: Query
}

type Query {
	people: [Character!]
}

type Character {
	name: String!
	height: String!
	mass: String!
	gender: String!
}
`

func (p Planet) Residents() []Character {
	var c []Character
	return c
}

var people AllPeople

func (r *Resolver) People() (*[]Character, error) {
	if len(people.People) == 0 {
		bytes, err := utils.LoadData("people")
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(bytes, &people); err != nil {
			return nil, err
		}
	}

	var s []Character

	for _, character := range people.People {
		s = append(s, character)
	}

	return &s, nil
}
