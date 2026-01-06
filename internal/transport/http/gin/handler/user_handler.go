package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
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

func (h *UserHandler) Create(c *gin.Context) {
	input := &entity.UserInput{
		FirstName: "john",
		LastName: "doe",
		Email: "0912student01@test.com",
		Password: "test1234",
	}
	
	result := h.userService.Create(c, input)
	if result != nil {
		c.JSON(http.StatusBadRequest, result)
	}

	c.JSON(http.StatusCreated, input)
}