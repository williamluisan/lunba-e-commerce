package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	users := rg.Group("/users")
	{
		// users.GET("/:id", h.GetUserById)
		users.GET("/email/:email", h.GetUserByEmail)
	}
}