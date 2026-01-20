package payment

import (
	"context"
)

type PaymentRepository interface {
	Process(ctx context.Context, orderId int64) (bool, error)
}