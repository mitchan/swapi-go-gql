package schema

import "github.com/mitchan/swapi-go-gql/types"

type Resolver struct{}

func (r *Resolver) People() (*[]types.Character, error) {
	var s []types.Character

	for _, character := range types.People.People {
		s = append(s, character)
	}

	return &s, nil
}

func (r *Resolver) Planets() (*[]types.Planet, error) {
	var slice []types.Planet

	for _, planet := range types.Planets.Planets {
		slice = append(slice, planet)
	}

	return &slice, nil
}

func (r *Resolver) Films() (*[]types.Film, error) {
	var slice []types.Film

	for _, planet := range types.Films.Films {
		slice = append(slice, planet)
	}

	return &slice, nil
}
