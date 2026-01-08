package repository

import (
	"context"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/repository"
	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &userRepository{db: db}
}

var mProductModel = gormModel.ProductModel{}

func (r *userRepository) Create(ctx context.Context, p *entity.Product) error {
	data := mProductModel.FromEntity(p)

	if err := gorm.G[gormModel.ProductModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return nil
}

func (r *userRepository) GetByCode(ctx context.Context, code string) (*entity.Product, error) {
	return nil, nil
}