package main

import (
	"log"

	transHttpGin "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/service"
	gormMysql "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"
	gormRepo "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/repository"
	productHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/product"
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
)

func main() {
	dsn := "lunba:toor@tcp(localhost:3306)/lunba_e_commerce?parseTime=true"
	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	/* infrastructure */
	// mysqlUser := mysql.NewMysqlUserRepository()
	gormUserRepo := gormRepo.NewUserRepository(db)
	gormProductRepo := gormRepo.NewProductRepository(db)

	/* domain */
	userRepo := gormUserRepo
	userService := service.NewUserService(userRepo)
	productRepo := gormProductRepo
	productService := service.NewProductService(productRepo)

	/* transport handler */
	userHandler := userHandler.NewUserHandler(userService)
	productHandler := productHandler.NewProductHandler(productService)

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		UserHandler: userHandler,
		ProductHandler: productHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":80")
}