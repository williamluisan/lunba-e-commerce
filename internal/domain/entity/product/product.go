package entity

import "time"

type Product struct {
	ID    		int64
	Name  		string
	Code  		string
	Price 		float64
	CreatedAt 	time.Time
	CreatedBy 	int
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
		CreatedBy: 1, // get from session
	}
}