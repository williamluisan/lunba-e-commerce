package gin

import (
	"github.com/gin-gonic/gin"

	"github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/routes"
)

func NewRouter(deps *Dependencies) *gin.Engine{
	r := gin.Default()

	api := r.Group("/api")

	routes.RegisterUserRoutes(api, deps.UserHandler)
	routes.RegisterProductRoutes(api, deps.ProductHandler)

	return r
}