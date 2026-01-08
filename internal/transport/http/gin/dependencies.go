package gin

import (
	productHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/product"
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	ProductHandler *productHandler.ProductHandler
}