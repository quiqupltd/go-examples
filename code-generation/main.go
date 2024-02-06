package main

import (
	"github.com/labstack/echo/v4"
	"github.com/quiqupltd/go-code-generation/api"
	"github.com/quiqupltd/go-code-generation/domain"
	"github.com/quiqupltd/go-code-generation/graph"
)

//go:generate go run github.com/99designs/gqlgen generate

//go:generate oapi-codegen -config oapi-codegen.yml ./spec/api.yml

func main() {
	e := echo.New()

	// Create a new domain service
	svc := domain.NewInMemoryProductsService()

	// Setup the GraphQL Server echo routes
	graph.ConfigGraphQLServer(e, svc)

	// Setup the REST APi echo routes
	err := api.ConfigureServer(e, svc)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
