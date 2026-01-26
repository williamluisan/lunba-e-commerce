package routes

import (
	productHandler "lunba-e-commerce/internal/transport/http/gin/handler/product"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, h *productHandler.ProductHandler) {
	product := rg.Group("/product") 
	{
		product.POST("/", h.Create)
	}
}