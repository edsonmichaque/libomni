package provider

import "io"

type ProviderBuilder func() Provider

type ConfigReader interface {
	ParseConfig(io.Reader) error
}

type Scheduler interface {
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

type Info interface {
	ID() string
	Name() string
	Version() string
}

type Provider interface {
	ConfigReader
	MessageSender
	MessageQuerier
	MessageReceiver
	Scheduler
	Capabilities
}