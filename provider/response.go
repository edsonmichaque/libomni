package provider

type Response struct {
	Error    *Error          `json:"error"`
	Messages []MessageStatus `json:"messages,omitempty"`
}
