package categoryservice

import (
	"context"

	"github.com/wellls/api-example-golang/internal/dto"
	"github.com/wellls/api-example-golang/internal/repository/categoryrepository"
)

func NewCategoryService(repo categoryrepository.CategoryRepository) CategoryService {
	return &service{
		repo,
	}
}

type service struct {
	repo categoryrepository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error
}
