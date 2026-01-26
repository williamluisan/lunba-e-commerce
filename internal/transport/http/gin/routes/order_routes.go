package routes

import (
	handler "lunba-e-commerce/internal/transport/http/gin/handler/order"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(rg *gin.RouterGroup, h *handler.OrderHandler) {
	order := rg.Group("/order")
	{
		order.POST("/", h.Create)
	}
}