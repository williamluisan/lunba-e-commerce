package order

import (
	"context"

	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/order"
	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/order"
)

type OrderService interface {
	Get(ctx context.Context, orderPublicId string) (*entity.Order, error)
	Create(ctx context.Context, e *entity.OrderInput) error
	Delete(ctx context.Context, orderPublicId string) error
}

type OrderServiceImpl struct {
	repo repository.OrderRepository
}

func New(r repository.OrderRepository) OrderService {
	if r == nil {
		panic("OrderService: repository not implemented.")
	}
	return &OrderServiceImpl{
		repo: r,
	}
}

func (i *OrderServiceImpl) Get(ctx context.Context, orderPublicId string) (*entity.Order, error) {
	return nil, nil
}

func (i *OrderServiceImpl) Create(ctx context.Context, e *entity.OrderInput) error {
	return nil
}

func (i *OrderServiceImpl) Delete(ctx context.Context, orderPublicId string) error {
	return nil
}

