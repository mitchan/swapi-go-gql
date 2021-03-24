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

func (c Character) Homeworld() *Planet {
	// search planet
	for _, planet := range planets.Planets {
		if planet.Url == c.HomeworldUrl {
			return &planet
		}
	}
	return nil
}

func (c Character) Films() *[]Film {
	var filmSlice []Film

	if len(c.FilmUrls) == 0 {
		return &filmSlice
	}

	for _, url := range c.FilmUrls {
		// search character
		for _, film := range films.Films {
			if film.Url == url {
				filmSlice = append(filmSlice, film)
				break
			}
		}
	}

	return &filmSlice
}

func (p Planet) Residents() *[]Character {
	var characters []Character

	if len(p.ResidentUrls) == 0 {
		return &characters
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

	return &characters
}

func (p Planet) Films() *[]Film {
	var filmSlice []Film

	if len(p.FilmUrls) == 0 {
		return &filmSlice
	}

	for _, url := range p.FilmUrls {
		// search character
		for _, film := range films.Films {
			if film.Url == url {
				filmSlice = append(filmSlice, film)
				break
			}
		}
	}

	return &filmSlice
}

func (f Film) Characters() *[]Character {
	var characters []Character

	if len(f.CharacterUrls) == 0 {
		return &characters
	}

	for _, url := range f.CharacterUrls {
		// search character
		for _, character := range people.People {
			if character.Url == url {
				characters = append(characters, character)
				break
			}
		}
	}

	return &characters
}

func (f Film) Planets() *[]Planet {
	var planetSlice []Planet

	if len(f.PlanetUrls) == 0 {
		return &planetSlice
	}

	for _, url := range f.PlanetUrls {
		// search planet
		for _, planet := range planets.Planets {
			if planet.Url == url {
				planetSlice = append(planetSlice, planet)
				break
			}
		}
	}

	return &planetSlice
}

var people AllPeople
var planets AllPlanets
var films AllFilms

var endpoints = []string{"people", "planets", "films"}

func PrefetchData() {
	for _, endpoint := range endpoints {
		switch endpoint {
		case "people":
			go loadEndpoint(endpoint, &people)

		case "planets":
			go loadEndpoint(endpoint, &planets)

		case "films":
			go loadEndpoint(endpoint, &films)
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

// func loadPeople() error {
// 	bytes, err := utils.LoadData("people")
// 	if err != nil {
// 		return err
// 	}

// 	if err := json.Unmarshal(bytes, &people); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func loadPlanets() error {
// 	bytes, err := utils.LoadData("planets")
// 	if err != nil {
// 		return err
// 	}

// 	if err := json.Unmarshal(bytes, &planets); err != nil {
// 		return err
// 	}

// 	return nil
// }

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

func (r *Resolver) Films() (*[]Film, error) {
	var slice []Film

	for _, planet := range films.Films {
		slice = append(slice, planet)
	}

	return &slice, nil
}
