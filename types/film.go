package types

type Film struct {
	Title         string   `json:"title"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	ReleaseDate   string   `json:"release_date"`
	Url           string   `json:"url"`
	CharacterUrls []string `json:"characters"`
	PlanetUrls    []string `json:"planets"`
}

type AllFilms struct {
	Films []Film `json:"results"`
}

var Films AllFilms

func (f Film) Characters() *[]Character {
	var characters []Character

	if len(f.CharacterUrls) == 0 {
		return &characters
	}

	for _, url := range f.CharacterUrls {
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

func (f Film) Planets() *[]Planet {
	var planetSlice []Planet

	if len(f.PlanetUrls) == 0 {
		return &planetSlice
	}

	for _, url := range f.PlanetUrls {
		// search planet
		for _, planet := range Planets.Planets {
			if planet.Url == url {
				planetSlice = append(planetSlice, planet)
				break
			}
		}
	}

	return &planetSlice
}
