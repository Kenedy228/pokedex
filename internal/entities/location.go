package entities

type Location struct {
	Areas []Area `json:"areas"`
}

type Area struct {
	Name string `json:"name"`
}
