package entity

//Dependency entity
type Dependency struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

//Health the health
type Health struct {
	Dependencies []Dependency `json:"dependencies"`
}
