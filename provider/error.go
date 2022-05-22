package provider

type ErrorCode string

const (
	ErrorNotImplemented = ErrorCode("NOT_IMPLEMENTED")
	ErrorNotSupported   = ErrorCode("NOT_SUPPORTED")
)

type Error struct {
	Code ErrorCode `json:"code,omitempty"`
}
