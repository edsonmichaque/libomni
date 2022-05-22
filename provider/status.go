package provider

type StatusCode string

const (
	StatusPending       = StatusCode("PENDING")
	StatusEnroute       = StatusCode("ENROUTE")
	StatusScheduled     = StatusCode("SCHEDULED")
	StatusDeleted       = StatusCode("DELETED")
	StatusAccepted      = StatusCode("ACCEPTED")
	StatusUnknown       = StatusCode("UNKNOWN")
	StatusSkipped       = StatusCode("SKIPPED")
	StatusUndeliverable = StatusCode("UNDELIVERABLE")
	StatusDelivered     = StatusCode("DELIVERED")
	StatusExpired       = StatusCode("EXPIRED")
	StatusRejected      = StatusCode("REJECTED")
)

type Status struct {
	Code        StatusCode `json:"code"`
	Reason      ReasonCode `json:"reason"`
	Description string     `json:"description"`
}
