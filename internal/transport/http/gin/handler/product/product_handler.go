package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/service"
	"github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductRequest

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

	input := &entity.ProductInput{
		Name: req.Name,
		Code: req.Code,
		Price: req.Price,
	}

	if err := h.productService.Create(c, input); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, handler.APIResponse{
		Success: true,
		Message: "Product created successfully",
	})
}