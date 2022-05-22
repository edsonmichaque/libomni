package provider

type MessageStatus struct {
	MessageID         string `json:"message_id,omitempty"`
	ProviderMessageID string `json:"provider_message_id,omitempty"`
	SubmittedAt       string `json:"submitted_at,omitempty"`
	DoneAt            string `json:"done_at,omitempty"`
	Status            Status `json:"status,omitempty"`
}
