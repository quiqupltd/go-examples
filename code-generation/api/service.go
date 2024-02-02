package api

import (
	"context"

	"github.com/quiqupltd/go-code-generation/domain"
)

type Service struct {
	svc domain.ProductsService
}

func NewService(svc domain.ProductsService) *Service {
	return &Service{svc: svc}
}

func (s *Service) GetAllProducts(ctx context.Context, request GetAllProductsRequestObject) (GetAllProductsResponseObject, error) {
	products, err := s.svc.Products()

	if err != nil {
		resp := GetAllProducts200JSONResponse{}
		return resp, err
	}

	resp := GetAllProducts200JSONResponse{}
	for _, p := range products {
		resp = append(resp, Product{
			Id:    p.ID,
			Name:  p.Name,
			Price: float32(p.Price),
		})
	}

	return resp, nil
}

func (s *Service) CreateProduct(ctx context.Context, request CreateProductRequestObject) (CreateProductResponseObject, error) {
	newProduct := domain.Product{
		ID:    request.Body.Id,
		Name:  request.Body.Name,
		Price: int64(request.Body.Price),
	}
	product, err := s.svc.CreateProduct(newProduct)
	if err != nil {
		resp := CreateProduct201JSONResponse{}
		return resp, err
	}

	resp := CreateProduct201JSONResponse{
		Id:    product.ID,
		Name:  product.Name,
		Price: float32(product.Price),
	}

	return resp, nil
}
