package greeter

import (
	"time"

	"github.com/quiqupltd/go-examples/wire/message"
)

type Greeter struct {
	Grumpy  bool
	Message message.Message
}

func (g Greeter) Greet() message.Message {
	if g.Grumpy {
		return message.Message("Go away!")
	}
	return g.Message
}

func NewGreeter(m message.Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}
