package main

import (
	"log"
	"time"

	"github.com/danhawkins/go-examples/functional-options/server"
)

func main() {
	svr := server.New(
		server.WithHost("localhost"),
		server.WithPort(8080),
		server.WithTimeout(time.Minute),
		server.WithMaxConn(120),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
