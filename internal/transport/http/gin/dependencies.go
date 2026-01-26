package gin

import (
	orderHandler "lunba-e-commerce/internal/transport/http/gin/handler/order"
	productHandler "lunba-e-commerce/internal/transport/http/gin/handler/product"
	userHandler "lunba-e-commerce/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	ProductHandler *productHandler.ProductHandler
	OrderHandler *orderHandler.OrderHandler
}