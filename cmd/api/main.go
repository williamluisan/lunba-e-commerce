package main

import (
	transHttpGin "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/service"
	"github.com/williamluisan/lunba-e-commerce/internal/infrastructure/mysql"
	"github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler"
)

func main() {
	// infrastructure
	mysqlUser := mysql.NewMysqlUserRepository()

	// domain
	userService := service.NewUserService(mysqlUser)

	// transport handler
	userHandler := handler.NewUserHandler(userService)

	// transport dependencies
	deps := &transHttpGin.Dependencies{
		UserHandler: userHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":80")
}