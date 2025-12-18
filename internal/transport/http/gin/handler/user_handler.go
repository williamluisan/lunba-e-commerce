package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	c.JSON(http.StatusOK, "user@gmail.com")
}