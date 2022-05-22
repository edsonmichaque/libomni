package provider

type Capability string

const (
	CapabilitySendMessage         = Capability("SEND_MESSAGE")
	CapabilityScheduleMessages    = Capability("SCHEDULE_MESSAGES")
	CapabilityQueryMessagesStatus = Capability("QUERY_MESSAGES")
	CapabilityReceiveMessages     = Capability("RECEIVE_MESSAGES")
)
