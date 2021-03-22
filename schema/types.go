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
