package event

import (
	"errors"
	"fmt"

	"github.com/danhawkins/go-examples/wire/greeter"
)

type Event struct {
	Greeter greeter.Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEvent(g greeter.Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
