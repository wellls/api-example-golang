package productservice

import (
	"context"

	"github.com/wellls/api-example-golang/internal/dto"
	"github.com/wellls/api-example-golang/internal/handler/response"
	"github.com/wellls/api-example-golang/internal/repository/productrepository"
)

func NewProductService(repo productrepository.ProductRepository) ProductService {
	return &service{
		repo,
	}
}

type service struct {
	repo productrepository.ProductRepository
}

type ProductService interface {
	CreateProduct(ctx context.Context, u dto.CreateProductDto) error
	UpdateProduct(ctx context.Context, id string, u dto.UpdateProductDto) error
	DeleteProduct(ctx context.Context, id string) error
	FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]response.ProductResponse, error)
}
