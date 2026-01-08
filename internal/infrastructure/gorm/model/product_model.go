package model

import (
	"time"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
)

type ProductModel struct {
	ID    		int64		`gorm:"primaryKey;unique;not null"`
	Name  		string		`gorm:"type:varchar(255);not null;"`
	Code  		string		`gorm:"type:varchar(50);not null;uniqueIndex"`
	Price 		float64		`gorm:"type:decimal(5,2);"`
	CreatedAt 	time.Time
	CreatedBy 	int
}

func (ProductModel) TableName() string {
	return "product"
}

func (m *ProductModel) FromEntity(p *entity.Product) *ProductModel {
	return &ProductModel{
		ID:    		p.ID,
		Name:  		p.Name,
		Code:  		p.Code,
		Price: 		p.Price,
		CreatedAt: 	p.CreatedAt,
		CreatedBy: 	p.CreatedBy,
	}
}

func (m *ProductModel) ToEntity() *entity.Product {
	return &entity.Product{
		ID:    		m.ID,
		Name:  		m.Name,
		Code:  		m.Code,
		Price: 		m.Price,
		CreatedAt: 	m.CreatedAt,
		CreatedBy: 	m.CreatedBy,
	}
} 