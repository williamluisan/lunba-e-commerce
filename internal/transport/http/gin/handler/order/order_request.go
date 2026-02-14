package order

type CreateOrderReq struct {
	ProductPublicId	string	`json:"product_pub_id" binding:"required"`
}