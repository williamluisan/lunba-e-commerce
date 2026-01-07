package gorm

import (
	"time"

	"github.com/google/uuid"
)

type ProductModel struct {
	ID			uuid.UUID	`gorm:"primaryKey;unique;not null"`
	Name		string		`gorm:"type:varchar(255); not null;"`
	Description string		`gorm:"type:varchar(255); not null;"`
	Code		string		`gorm:"type:varchar(255); not null;"`
	Price		float64		`gorm:"type:decimal(5,2); not null;"`
	CreatedAt	time.Time
	CreatedBy	int
}

func (ProductModel) TableName() string {
	return "products"
}