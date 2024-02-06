package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/quiqupltd/go-code-generation/domain"
)

func ConfigGraphQLServer(server *echo.Echo, svc domain.ProductsService) error {
	// Create a new graphql handler
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{ProductsService: svc}}))

	// Setup the graphql playground
	server.GET("/play", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))

	// Connect our handler at /query
	server.POST("/query", echo.WrapHandler(srv))

	return nil
}
