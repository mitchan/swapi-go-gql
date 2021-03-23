package schema

type Character struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
	Gender string `json:"gender"`
	Url    string `json:"url"`
}

type AllPeople struct {
	People []Character `json:"results"`
}

type Planet struct {
	Name         string   `json:"name"`
	Climate      string   `json:"climate"`
	Terrain      string   `json:"terrain"`
	Population   string   `json:"population"`
	ResidentUrls []string `json:"residents"`
}

type AllPlanets struct {
	Planets []Planet `json:"results"`
}

type Resolver struct{}
