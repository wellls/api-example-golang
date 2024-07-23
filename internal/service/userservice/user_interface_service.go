package userservice

import (
	userrepository "github.com/wellls/api-example-golang/internal/repository/userepository"
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
	CreateUser() error
}
