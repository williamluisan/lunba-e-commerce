package repository

import (
	"context"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/repository"
	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) Create(ctx context.Context, u *entity.User) error {
	var m gormModel.UserModel

	data := m.FromEntity(u)

	if err := gorm.G[gormModel.UserModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, u *entity.User) error {
	return nil
}