package routes

import (
	userHandler "lunba-e-commerce/internal/transport/http/gin/handler/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *userHandler.UserHandler) {
	user := rg.Group("/user")
	{
		user.POST("/", h.Create)
		// user.GET("/:id", h.GetUserById)
		user.GET("/email/:email", h.GetUserByEmail)
	}
}