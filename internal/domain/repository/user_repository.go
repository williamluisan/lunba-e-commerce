package repository

import (
	"context"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, u *entity.UserInput) error
	Update(ctx context.Context, u *entity.UserInput) error
}