# Go Code Generation Examples

## Introduction

Code generation in Go (Golang) is a technique where you write programs that generate other programs, or parts of a program. This is often used to automate repetitive tasks, reduce human error, and maintain consistency.

### Pros of Code Generation in Go

* Automation: Code generation can automate repetitive tasks, reducing the chance of human error and increasing efficiency.
* Consistency: It ensures consistency across different parts of the codebase. If you have a common pattern that is used in many places, generating the code for that pattern ensures it is the same everywhere.
* Performance: In some cases, generating code can lead to more efficient programs. For example, generating specific versions of a function for different types can be more efficient than a generic version that works with interfaces.
* Focused Refactoring / Change management: When changing the API and regenerating code, areas that need fixing fail compilation, and are picked up by go language server, giving a precise list problems in the codebase.

### Cons of Code Generation in Go

* Complexity: Code generation can add complexity to the codebase. It's another thing that developers need to understand and maintain.
* Debugging: Debugging generated code can be difficult because the code that you see in the debugger is not the code that you wrote.
* Readability: Generated code can be harder to read and understand than hand-written code.

In Go, code generation is often used for tasks like generating String or JSON methods for types, generating database access code from schema definitions, creating boilerplate code or generating specific versions of functions for different types. The [go generate](https://go.dev/blog/generate) command is a tool that's built into the Go toolchain to help with code generation.

## Running the server

From the code-generation folder run `go mod tidy`, then `go run .` the server will be available at http://localhost:8080,

* [GraphQL Playground](https://localhost:8080/play)
* [Swagger UI](https://localhost:8080/swaggerui)

### Example Graphql Queries

#### Mutation

```graphql
mutation CreateProduct($product:ProductInput!) {
  createProduct(input: $product) {
    id
    name
    price
  }
}
```

#### Variables

```json
{
  "product": {
    "name": "Testing 1 2 3",
    "price": 500
  }
}
```

#### Query

```graphql
query {
  getAllProducts {
    id
    name
    price
  }
}
```

## Goals

In this example we have an [openapi specification](./spec/api.yml) and a [graphql specifcation](./spec/graph.gql) we want to generate the code that provides implementation poits for the protocols, without having to write specifics for the protocol (at least avoid as much as possible).

> **IMPORTANT: Any files that are generated will state such at the top of the file, and these files are not expected to be editing by the developer**

### Starting point

We start with two specification files, and a [domain service](domain/service.go) under the domain module, this domain has a single service which can list products and create products in the list using an in memory array. This will service as the core of our app, or in hexagonal architecture this is at the centre.

## Open APi code gen

In order to generate the server code to implement the openapi spec we use [oapi-codegen](https://github.com/deepmap/oapi-codegen) which is a go library to generate servers and clients, there is a more widely recognised code generator with more options called [openapi-generator](https://github.com/OpenAPITools/openapi-generator) however the former supports go server generation in a much better way, especially for the framework we will use called [echo](https://echo.labstack.com/).

To use the generator we add a [config file](./oapi-codegen.yml) and a generate comment in [main.go](./main.go#L12)

```go
//go:generate oapi-codegen -config oapi-codegen.yml ./spec/api.yml
```

Running `go generate ./...` will then create the [output file](./api/open_api.gen.go)

Now we need to create implementation code to fullfil the interface defined in the generated code. There are two choices normal or strict interfaces (this is configurable via the config file, and you can read more about the options on the github repo)

Normal:

```go
type ServerInterface interface {
	// Get all products
	// (GET /products)
	GetAllProducts(ctx echo.Context) error
	// Create a new product
	// (POST /products)
	CreateProduct(ctx echo.Context) error
}
```

Strict:

```go
type StrictServerInterface interface {
	// Get all products
	// (GET /products)
	GetAllProducts(ctx context.Context, request GetAllProductsRequestObject) (GetAllProductsResponseObject, error)
	// Create a new product
	// (POST /products)
	CreateProduct(ctx context.Context, request CreateProductRequestObject) (CreateProductResponseObject, error)
}
```

I prefer to use the strict interface, as it's much harder to break a contract, and the IDE gives you much more feedback on if things are wront, at the cost of it being a little bit more complex to understand and implement.

We create the [implementation service](./api/service.go) to connect our API to our [domain service](./domain/service.go).

All of the REST API and Swagger routing for echo is configuted in [the server file](./api/server.go) and added to the echo server in [main.go](./main.go#L24)

Theres a few more things going on in [server.go](./api/server.go), I've added inline comments to that file to explain details.

## GraphQL code gen

For the GraphQL codegen we use [gqlgen](https://github.com/99designs/gqlgen), the principal is the same as openapi codegen, we have a [spec file](./spec/graph.gql) and a [config file](./gqlgen.yml) altough the config file has many more options, refer to the [documentation site](https://gqlgen.com/getting-started/) for more details.

Then we add the generator line in [main.go](./main.go#L10), running `go generate ./...` will then generate the graphql code according to the spec.

```go
//go:generate oapi-codegen -config oapi-codegen.yml ./spec/api.yml
```

Implementation of the service is little different with the gql generated code. First of all you will need to add any dependencies from your implementation to [resolver.go](./graph/resolver.go) this is just a plain struct the you will need to initialise. You can see this in [server.go](./graph/server.go#L12).

```go
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{ProductsService: svc}}))
```

The [graph.resolvers.go](./graph/graph.resolvers.go) is where you add the implementation, the generate will take care of adding new functions there, or getting rid of old ones as the spec changes. This file is generated, but in this case you are expected to change the file adding your implementation (as stated in the comment at the top of the file)

We also setup a graphiql playground at `/play` in [server.go](./graph/server.go#L15)
