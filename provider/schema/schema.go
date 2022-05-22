package schema

import "encoding/json"

type Properties map[string]Schema

type Schema struct {
	Type              kind       `json:"type,omitempty"`
	Format            format     `json:"format,omitempty"`
	Description       string     `json:"description,omitempty"`
	Properties        Properties `json:"properties,omitempty"`
	Required          []string   `json:"required,omitempty"`
	Items             *Schema    `json:"items,omitempty"`
	PatternProperties Properties `json:"patternProperties,omitempty"`
	Secret            bool       `json:"secret,omitempty"`
}

type kind int

const (
	String kind = iota
	Number
	Object
	Boolean
	Array
)

func (k kind) String() string {
	return map[kind]string{
		String:  "string",
		Number:  "number",
		Object:  "object",
		Boolean: "boolean",
		Array:   "array",
	}[k]
}

func (k kind) MarshalJSON() ([]byte, error) {
	return []byte(k.String()), nil
}

func (k kind) UmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	kinds := map[string]kind{
		"string":  String,
		"number":  Number,
		"object":  Object,
		"boolean": Boolean,
		"array":   Array,
	}

	if ks, ok := kinds[value]; ok {
		*(&k) = ks
	}

	return nil
}

func (k kind) Validate() error {
	return nil
}

type format int

const (
	Email format = iota
)
