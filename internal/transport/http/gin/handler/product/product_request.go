package product

type CreateProductRequest struct {
	Name 	string 	`form:"name" json:"name" binding:"required"`
	Code	string	`form:"code" json:"code" binding:"required"`
	Price	float64 `form:"price" json:"price" binding:"required"`
}