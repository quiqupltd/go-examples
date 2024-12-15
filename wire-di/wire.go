//go:build wireinject
// +build wireinject

package main

import (
	"github.com/danhawkins/go-examples/wire/event"
	"github.com/danhawkins/go-examples/wire/greeter"
	"github.com/danhawkins/go-examples/wire/message"
	"github.com/google/wire"
)

func InitializeEvent(phrase string) (event.Event, error) {
	wire.Build(event.NewEvent, greeter.NewGreeter, message.NewMessage)
	return event.Event{}, nil
}
