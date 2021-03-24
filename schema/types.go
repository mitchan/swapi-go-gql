package schema

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

type Resolver struct{}
