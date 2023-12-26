package main

import (
	"fmt"

	"github.com/rs/xid"
)

type ResourceType interface{ Prefix() string }
type ID[T ResourceType] xid.ID

type Mission struct{}

func (u Mission) Prefix() string { return "mission" }

type Order struct{}

func (a Order) Prefix() string { return "order" }

func NewID[T ResourceType]() ID[T] { return ID[T](xid.New()) }

func (id ID[T]) String() string {
	var resourceType T // create the default value for the resource type

	return fmt.Sprintf(
		"%s_%s",
		resourceType.Prefix(), // Extract the "prefix" we want from the resource type
		xid.ID(id).String(),   // Use XID's string marshalling
	)
}

func main() {
	fmt.Println(NewID[Mission]())

	fmt.Println(NewID[Order]())
}
