package order

import (
	entity "lunba-e-commerce/internal/domain/entity/order"
	repository "lunba-e-commerce/internal/domain/repository/order"
	gormModel "lunba-e-commerce/internal/infrastructure/gorm/model"

	"context"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository.OrderRepository {
	if db == nil {
		panic("Gorm OrderRepository: didn't pass db instance")
	}
	return &orderRepository{db: db}
}

var mOrderModel = gormModel.OrderModel{}

func (r *orderRepository) Get(ctx context.Context, orderPublicId string) (*entity.Order, error) {
	return nil, nil
}

func (r *orderRepository) GetByProductAndUser(ctx context.Context, productPublicId string, userPublicId string) (*entity.Order, error) {
	// result, err := nil, nil

	return nil, nil
}

func (r *orderRepository) Create(ctx context.Context, e *entity.Order) error {
	data := mOrderModel.FromEntity(e)

	if err := gorm.G[gormModel.OrderModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) Delete(ctx context.Context, orderPublicId string) error {
	return nil
}
