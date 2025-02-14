//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/quiqupltd/go-examples/wire/event"
	"github.com/quiqupltd/go-examples/wire/greeter"
	"github.com/quiqupltd/go-examples/wire/message"
)

func InitializeEvent(phrase string) (event.Event, error) {
	wire.Build(event.NewEvent, greeter.NewGreeter, message.NewMessage)
	return event.Event{}, nil
}
