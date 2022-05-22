package provider

import "io"

type Provider interface {
	ParseConfig(io.Reader) error
	SendMessages(Request) Response
	QueryMessagesStatus(Request) Response
	ReceiveMessages(Request) Response
	ScheduleMessages(string, Request) Response
	Capabilities() []Capability
}
