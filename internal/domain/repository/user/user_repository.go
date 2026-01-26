package repository

import (
	"context"

	entity "lunba-e-commerce/internal/domain/entity/user"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, u *entity.User) error
	Update(ctx context.Context, u *entity.User) error
}