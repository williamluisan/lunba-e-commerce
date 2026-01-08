package model

import "time"

type ProductModel struct {
	ID    		int64		`gorm:"primaryKey;unique;not null"`
	Name  		string		`gorm:"type:varchar(255);not null;"`
	Code  		string		`gorm:"type:varchar(50);not null;"`
	Price 		float64		`gorm:"type:decimal(5,2);"`
	CreatedAt 	time.Time
	CreatedBy 	int
}

func (ProductModel) TableName() string {
	return "product"
}