package order

import (
	"context"

	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/order"
)

type OrderRepository interface {
	Get(ctx context.Context, orderPublicId string) (*entity.Order, error)
	Create(ctx context.Context, e *entity.Order) error
	Delete(ctx context.Context, orderPublicId string) error
}