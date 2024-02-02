package domain

// Product for the internal domain
type Product struct {
	ID    *string `json:"id"`    // ID
	Name  string  `json:"name"`  // Name of product
	Price int64   `json:"price"` // Price in cents
}

// ProductsService for the internal domain interface
type ProductsService interface {
	Products() ([]*Product, error)
	CreateProduct(input Product) (*Product, error)
}

// In memory implementation of the ProductsService
type productsService struct {
	products []*Product
}

func NewInMemoryProductsService() ProductsService {
	return &productsService{
		products: []*Product{},
	}
}

// List all products
func (s *productsService) Products() ([]*Product, error) {
	return s.products, nil
}

// CreateProduct creates a new product in the list
func (s *productsService) CreateProduct(input Product) (*Product, error) {
	s.products = append(s.products, &input)
	return &input, nil
}
