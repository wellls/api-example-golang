package userservice

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/wellls/api-example-golang/internal/dto"
	"github.com/wellls/api-example-golang/internal/entity"
	"github.com/wellls/api-example-golang/internal/handler/response"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
		return err
	}
	if userExists != nil {
		slog.Error("user already exists", slog.String("package", "userservice"))
		return errors.New("user already exists")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}

	newUser := entity.UserEntity{
		ID:        uuid.New().String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  string(passwordEncrypted),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.CreateUser(ctx, &newUser)
	if err != nil {
		slog.Error("error to create user", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
	userFake := response.UserResponse{
		ID:        "123",
		Name:      "John Doe",
		Email:     "jonh.doe@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &userFake, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (s *service) FindManyUsers(ctx context.Context) (response.ManyUsersResponse, error) {
	// create fake users
	usersFake := response.ManyUsersResponse{}
	for i := 0; i < 5; i++ {
		userFake := response.UserResponse{
			ID:        "123",
			Name:      "John Doe",
			Email:     fmt.Sprintf("jonh.doe-%v@email.com", i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		usersFake.Users = append(usersFake.Users, userFake)
	}
	return usersFake, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
	fmt.Println("new password: ", u.Password)
	fmt.Println("old password: ", u.OldPassword)
	return nil
}
