package graph

import "github.com/quiqupltd/go-code-generation/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductsService domain.ProductsService
}
