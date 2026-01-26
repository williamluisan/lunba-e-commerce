package product

import (
	"context"

	entity "lunba-e-commerce/internal/domain/entity/product"
	repository "lunba-e-commerce/internal/domain/repository/product"
	gormModel "lunba-e-commerce/internal/infrastructure/gorm/model"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

var mProductModel = gormModel.ProductModel{}

func (r *productRepository) Create(ctx context.Context, p *entity.Product) error {
	data := mProductModel.FromEntity(p)

	if err := gorm.G[gormModel.ProductModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	return nil
}

func (r *productRepository) GetByCode(ctx context.Context, code string) (*entity.Product, error) {
	return nil, nil
}