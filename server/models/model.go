package models

// CharShit .
type CharShit struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Core interface{} `json:"core"`
}

// NewChar .
func NewChar(name string) CharShit {
	return CharShit{
		Name: name,
	}
}
