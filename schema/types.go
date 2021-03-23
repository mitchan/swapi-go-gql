package schema

type Character struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
	Gender string `json:"gender"`
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

type characterResolver struct {
	name string
}

type planetResult struct {
	result interface{}
}

// RESOLVER
type Resolver struct{}
