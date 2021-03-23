package schema

import (
	"encoding/json"

	"github.com/mitchan/swapi-go-gql/utils"
)

var Schema = `
schema {
	query: Query
}

type Query {
	people: [Character!]
	planets: [Planet!]
}

type Character {
	name: String!
	height: String!
	mass: String!
	gender: String!
}

type Planet {
	name: String!
	climate: String!
	terrain: String!
	population: String!
	residents: [Character!]
}
`

func (p Planet) Residents() (*[]Character, error) {
	var characters []Character

	if len(p.ResidentUrls) == 0 {
		return &characters, nil
	}

	if len(people.People) == 0 {
		if err := loadPeople(); err != nil {
			return nil, err
		}
	}

	for _, url := range p.ResidentUrls {
		// search character
		for _, character := range people.People {
			if character.Url == url {
				characters = append(characters, character)
				break
			}
		}
	}

	return &characters, nil
}

var people AllPeople
var planets AllPlanets

var endpoints = []string{"people", "planets"}

func PrefetchData() {
	for _, endpoint := range endpoints {
		switch endpoint {
		case "people":
			go loadPeople()

		case "planets":
			go loadPlanets()
		}
	}
}

func loadPeople() error {
	bytes, err := utils.LoadData("people")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &people); err != nil {
		return err
	}

	return nil
}

func loadPlanets() error {
	bytes, err := utils.LoadData("planets")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &planets); err != nil {
		return err
	}

	return nil
}

func (r *Resolver) People() (*[]Character, error) {
	var s []Character

	for _, character := range people.People {
		s = append(s, character)
	}

	return &s, nil
}

func (r *Resolver) Planets() (*[]Planet, error) {
	var slice []Planet

	for _, planet := range planets.Planets {
		slice = append(slice, planet)
	}

	return &slice, nil
}
