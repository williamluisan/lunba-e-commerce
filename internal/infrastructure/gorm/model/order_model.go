package model

import (
	entity "lunba-e-commerce/internal/domain/entity/order"
	"time"
)

type OrderModel struct {
	ID				uint64			`gorm:"primaryKey;unique;not null"`
	PublicId		string			`gorm:"type:char(26);not null"`
	UserPublicId	string			`gorm:"type:char(26);not null"`
	ProductPublicId	string			`gorm:"type:char(26);not null"`
	IsPaid			bool			`gorm:"type:int(1);not null;default:0"`
	Qty				int				`gorm:"type:int(11);not null;default:1"`
	CreatedAt		time.Time		`gorm:"not null"`
	CreatedBy		string			`gorm:"not null;type:char(26)"`
	UpdatedAt		time.Time
	UpdatedBy		string			`gorm:"type:char(26)"`
}

func (OrderModel) TableName() string {
	return "orders"
}

func (m *OrderModel) FromEntity(e *entity.Order) *OrderModel {
	return &OrderModel{
		ID:	e.ID,
		PublicId: e.PublicId,
		UserPublicId: e.UserPublicId,
		ProductPublicId: e.ProductPublicId,
	}
}

func (m *OrderModel) ToEntity() *entity.Order {
	return &entity.Order{
		ID:	m.ID,
		PublicId: m.PublicId,
		UserPublicId: m.UserPublicId,
		ProductPublicId:m.ProductPublicId,
		IsPaid: m.IsPaid,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
	}
}