package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID			uuid.UUID
	OrderId		uint64
	Amount		float64
	Status		string
	CreatedAt	time.Time
	CreatedBy	int
}