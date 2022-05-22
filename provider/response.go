package provider

var (
	ResponseNotImplemented = Response{
		Error: &Error{
			Code: ErrorNotImplemented,
		},
	}

	ResponseNotSupported = Response{
		Error: &Error{
			Code: ErrorNotSupported,
		},
	}
)

type Response struct {
	Error    *Error          `json:"error"`
	Messages []MessageStatus `json:"messages,omitempty"`
}
