package schema

type Schema struct {
	Type              Kind       `json:"type"`
	Format            Format     `json:"format"`
	Description       string     `json:"description"`
	Properties        Properties `json:"properties"`
	Required          []string   `json:"required"`
	Items             []Schema   `json:"items"`
	PatternProperties Properties `json:"patternProperties"`
}

type Kind string

const (
	String  = Kind("string")
	Number  = Kind("number")
	Object  = Kind("object")
	Boolean = Kind("boolean")
	Array   = Kind("array")
)

type Format string

const (
	Email = Format("email")
)

func (k Kind) Validate() error {
	return nil
}

type Properties map[string]Schema
