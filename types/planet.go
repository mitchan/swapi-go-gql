package types

type Planet struct {
	Name         string   `json:"name"`
	Climate      string   `json:"climate"`
	Terrain      string   `json:"terrain"`
	Population   string   `json:"population"`
	Url          string   `json:"url"`
	ResidentUrls []string `json:"residents"`
	FilmUrls     []string `json:"films"`
}

type AllPlanets struct {
	Planets []Planet `json:"results"`
}

var Planets AllPlanets

func (p Planet) Residents() *[]Character {
	var characters []Character

	if len(p.ResidentUrls) == 0 {
		return &characters
	}

	for _, url := range p.ResidentUrls {
		// search character
		for _, character := range People.People {
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
		for _, film := range Films.Films {
			if film.Url == url {
				filmSlice = append(filmSlice, film)
				break
			}
		}
	}

	return &filmSlice
}
