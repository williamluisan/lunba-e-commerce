package order

type CreateOrderReq struct {
	UserPublicId	string	`json:"user_pub_id" binding:"required"` // TODO: this need to get from the key
	ProductPublicId	string	`json:"product_pub_id" binding:"required"`
}