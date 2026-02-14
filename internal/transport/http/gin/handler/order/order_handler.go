package order

import (
	"errors"
	"net/http"

	entity "lunba-e-commerce/internal/domain/entity/order"
	messages "lunba-e-commerce/internal/domain/messages"
	"lunba-e-commerce/internal/transport/http/gin/handler"
	service "lunba-e-commerce/internal/usecase/order"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service service.OrderService
}

func New(s service.OrderService) *OrderHandler {
	if s == nil {
		panic("OrderHandler (GIN): order service not implemented.")
	}
	return &OrderHandler{
		service: s,
	}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req CreateOrderReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: messages.ValidationError.Message,
			Error: &handler.APIError{
				Code: messages.ValidationError.Code,
				Message: messages.ValidationError.Message,
				Details: handler.ParseValidationError(err),
			},
		})
		return
	}

	input := &entity.OrderInput{
		ProductPublicId: req.ProductPublicId,
	}

	if err := h.service.Create(c, input); err != nil {
		var e *messages.OrderAlreadyExistsError
		if errors.As(err, &e) {
			c.JSON(http.StatusConflict, handler.APIResponse{
				Success: false,
				Message: e.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, handler.APIResponse{
		Success: true,
		Message: messages.OrderCreated.Message,
	})
}