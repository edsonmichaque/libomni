package provider

type Request struct {
	Messages []Message `json:"messages,omitempty"`
	Schedule *string   `json:"schedule"`
}
