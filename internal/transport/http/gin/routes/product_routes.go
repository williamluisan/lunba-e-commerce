package routes

import (
	"github.com/gin-gonic/gin"
	productHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/product"
)

func RegisterProductRoutes(rg *gin.RouterGroup, h *productHandler.ProductHandler) {
	product := rg.Group("/product") 
	{
		product.POST("/", h.Create)
	}
}