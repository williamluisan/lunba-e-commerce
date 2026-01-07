package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *userHandler.UserHandler) {
	users := rg.Group("/users")
	{
		users.POST("/create", h.Create)
		// users.GET("/:id", h.GetUserById)
		users.GET("/email/:email", h.GetUserByEmail)
	}
}