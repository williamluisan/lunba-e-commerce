package entity

import "time"

type Product struct {
	ID    		int64
	PublicId 	string
	Name  		string
	Code  		string
	Price 		float64
	Stock		int
	Status		string
	CreatedAt 	time.Time
	CreatedBy 	int
	UpdatedAt	time.Time
	UpdatedBy	int
}

// DTO
type ProductInput struct {
	Name  		string
	Code  		string
	Price 		float64
}

func NewProduct(input *ProductInput) *Product {
	return &Product{
		Name: input.Name,
		Code: input.Code,
		Price: input.Price,
		CreatedAt: time.Now(),
		CreatedBy: 1, // TODO: get from token/session
	}
}