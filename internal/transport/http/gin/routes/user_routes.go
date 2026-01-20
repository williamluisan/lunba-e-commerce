package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *userHandler.UserHandler) {
	user := rg.Group("/user")
	{
		user.POST("/", h.Create)
		// user.GET("/:id", h.GetUserById)
		user.GET("/email/:email", h.GetUserByEmail)
	}
}