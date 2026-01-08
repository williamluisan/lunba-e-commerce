package repository

import (
	"context"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, p *entity.Product) error
	Delete(ctx context.Context, id int64) error
	GetByCode(ctx context.Context, code string) (*entity.Product, error) 
}