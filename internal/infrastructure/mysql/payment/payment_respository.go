package payment

import (
	"context"

	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/payment"
)

type paymentRepositoryImpl struct {}

func New() repository.PaymentRepository {
	return &paymentRepositoryImpl{}
}

func (r *paymentRepositoryImpl) Process(ctx context.Context, orderId int64) (bool, error) {
	return true, nil
} 