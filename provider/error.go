package provider

type ErrorType string

const (
	ErrorNotImplemented = ErrorType("NOT_IMPLEMENTED")
	ErrorNotSupported   = ErrorType("NOT_SUPPORTED")
	ErrorConnection     = ErrorType("CONNECTION")
	ErrorUnknown        = ErrorType("UNKNOWN")
)

type Error struct {
	Code string    `json:"code"`
	Type ErrorType `json:"type"`
	Desc string    `json:"desc"`
}

var errorCodes = map[ErrorType]string{
	ErrorNotImplemented: "10",
	ErrorNotSupported:   "11",
	ErrorConnection:     "12",
	ErrorUnknown:        "13",
}
