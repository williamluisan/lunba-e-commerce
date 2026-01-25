package model

import "time"

type OrderModel struct {
	ID				uint64		`gorm:"primaryKey;unique;not null"`
	PublicId		string		`gorm:"type:char(26);not null"`
	UserPublicId	string		`gorm:"type:char(26);not null"`
	ProductPublicId	string		`gorm:"type:char(26);not null"`
	IsPaid			bool		`gorm:"type:int(1);not null;default:0"`
	CreatedAt		time.Time	`gorm:"not null"`
	CreatedBy		int			`gorm:"not null;type:int(10)"`
	UpdatedAt		time.Time
	UpdatedBy		int			`gorm:"type:int(10)"`
}

func (OrderModel) TableName() string {
	return "orders"
}