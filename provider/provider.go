package provider

import "io"

type Builder func() Provider

type ConfigReader interface {
	ParseConfig(io.Reader) error
}

type MessageScheduler interface {
	ScheduleMessages(string, Request) Response
}

type MessageSender interface {
	SendMessages(Request) Response
}

type MessageQuerier interface {
	QueryMessagesStatus(Request) Response
}

type MessageReceiver interface {
	ReceiveMessages(Request) Response
}

type Capabilities interface {
	Capabilities() []Capability
}

type Provider interface {
	ConfigReader
	MessageSender
	MessageQuerier
	MessageReceiver
	MessageScheduler
	Capabilities
}
