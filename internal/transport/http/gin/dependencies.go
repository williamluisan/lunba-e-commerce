package gin

import (
	orderHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/order"
	productHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/product"
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	ProductHandler *productHandler.ProductHandler
	OrderHandler *orderHandler.OrderHandler
}