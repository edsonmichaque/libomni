package provider

type Message struct {
	ID           string        `json:"id,omitempty"`
	ProviderID   string        `json:"provider_id,omitempty"`
	Text         string        `json:"text,omitempty"`
	From         string        `json:"from,omitempty"`
	SubmittedAt  string        `json:"submitted_at,omitempty"`
	DoneAt       string        `json:"done_at,omitempty"`
	Destinations []Destination `json:"destinations,omitempty"`
	Status       Status        `json:"status,omitempty"`
	Quotation    Quotation     `json:"quotation,omitempty"`
}
