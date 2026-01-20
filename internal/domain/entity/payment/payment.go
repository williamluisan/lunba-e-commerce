package payment

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID			uuid.UUID
	OrderId		int64
	Amount		float64
	Status		string
	CreatedAt	time.Time
	CreatedBy	int
}

type PaymentProcess struct {
	OrderId				int64
	OutstandingAmount	float64
	PaidAmount			float64
	Balance				float64
	RandomAmount		float64
}