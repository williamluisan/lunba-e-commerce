package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/order"
	"github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler"
	service "github.com/williamluisan/lunba-e-commerce/internal/usecase/order"
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
			Error: &handler.APIError{
				Code: "validation_error",
				Message: "Validation error",
				Details: handler.ParseValidationError(err),
			},
		})
		return
	}

	// check if order public id is exists (on laravel microservice)
	// ...

	// check if product is exists via public id (on laravel microservices)
	// ...

	input := &entity.OrderInput{
		UserPublicId: req.UserPublicId,
		ProductPublicId: req.ProductPublicId,
	}

	if err := h.service.Create(c, input); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, handler.APIResponse{
		Success: true,
		Message: "Order created successfully",
	})
}