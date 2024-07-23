package userservice

import (
	"context"

	"github.com/wellls/api-example-golang/internal/dto"
	"github.com/wellls/api-example-golang/internal/repository/userrepository"
)

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
}
