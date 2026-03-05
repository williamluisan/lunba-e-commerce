package repository

import (
	"context"

	entity "lunba-e-commerce/internal/domain/entity/product"
)

type ProductRepository interface {
	Create(ctx context.Context, p *entity.Product) error
	Delete(ctx context.Context, id int64) error
	GetByCode(ctx context.Context, code string) (*entity.Product, error) 
}

// external integration repository interface for product
type ProductRepositoryExt interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error)
}