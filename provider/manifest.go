package provider

type Manifest struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Desc    string `json:"desc"`
	Schema  string `json:"schema"`
}
