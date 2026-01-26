package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/order"
)

func RegisterOrderRoutes(rg *gin.RouterGroup, h *handler.OrderHandler) {
	order := rg.Group("/order")
	{
		order.POST("/", h.Create)
	}
}