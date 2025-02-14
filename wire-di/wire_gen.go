// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/quiqupltd/go-examples/wire/event"
	"github.com/quiqupltd/go-examples/wire/greeter"
	"github.com/quiqupltd/go-examples/wire/message"
)

// Injectors from wire.go:

func InitializeEvent(phrase string) (event.Event, error) {
	messageMessage := message.NewMessage(phrase)
	greeterGreeter := greeter.NewGreeter(messageMessage)
	eventEvent, err := event.NewEvent(greeterGreeter)
	if err != nil {
		return event.Event{}, err
	}
	return eventEvent, nil
}
