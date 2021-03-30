package types

type Character struct {
	Name         string   `json:"name"`
	Height       string   `json:"height"`
	Mass         string   `json:"mass"`
	Gender       string   `json:"gender"`
	HomeworldUrl string   `json:"homeworld"`
	Url          string   `json:"url"`
	FilmUrls     []string `json:"films"`
}

type AllPeople struct {
	People []Character `json:"results"`
}

var People AllPeople

func (c Character) Homeworld() *Planet {
	// search planet
	for _, planet := range Planets.Planets {
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
		for _, film := range Films.Films {
			if film.Url == url {
				filmSlice = append(filmSlice, film)
				break
			}
		}
	}

	return &filmSlice
}
