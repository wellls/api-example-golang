package userrepository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/wellls/api-example-golang/internal/database/sqlc"
	"github.com/wellls/api-example-golang/internal/entity"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
	err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})
	if err != nil {
		return err
	}
	err = r.queries.CreateUserAddress(ctx, sqlc.CreateUserAddressParams{
		ID:         uuid.New().String(),
		UserID:     u.ID,
		Cep:        u.Address.CEP,
		Ibge:       u.Address.IBGE,
		Uf:         u.Address.UF,
		City:       u.Address.City,
		Complement: sql.NullString{String: u.Address.Complement, Valid: u.Address.Complement != ""},
		Street:     u.Address.Street,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	userMock := entity.UserEntity{
		ID:    "1",
		Name:  "John",
		Email: "john.doe@email.com",
	}
	return &userMock, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	return nil
}

func (r *repository) GetUserPassword(ctx context.Context, id string) (*entity.UserEntity, error) {
	userMock := entity.UserEntity{
		ID:       "1",
		Password: "$2y$12$CwjjXJGAkR4OKQeTvMo9suJ1s6PdKl9l4RZL9/yg.8cccDE8o/5sm",
	}
	return &userMock, nil
}
