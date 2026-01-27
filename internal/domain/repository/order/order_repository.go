package order

import (
	"context"

	entity "lunba-e-commerce/internal/domain/entity/order"
)

type OrderRepository interface {
	Get(ctx context.Context, orderPublicId string) (*entity.Order, error)
	GetByProductAndUser(ctx context.Context, productPublicId string, userPublicId string) (*entity.Order, error)
	Create(ctx context.Context, e *entity.Order) error
	Delete(ctx context.Context, orderPublicId string) error
}