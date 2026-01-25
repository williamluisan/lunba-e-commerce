package main

import (
	"log"

	transHttpGin "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin"

	gormMysql "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"
	gormProductRepo "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/repository/product"
	gormUserRepo "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/repository/user"
	productHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/product"
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
	productUsecase "github.com/williamluisan/lunba-e-commerce/internal/usecase/product"
	userUsecase "github.com/williamluisan/lunba-e-commerce/internal/usecase/user"
)

func main() {
	dsn := "root:toor@tcp(mysql:3306)/ln-e-commerce_order_payment?parseTime=true"
	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	/* infrastructure */
	// mysqlUser := mysql.NewMysqlUserRepository()
	gormUserRepo := gormUserRepo.NewUserRepository(db)
	gormProductRepo := gormProductRepo.NewProductRepository(db)

	/* domain */
	userRepo := gormUserRepo
	userService := userUsecase.NewUserService(userRepo)
	productRepo := gormProductRepo
	productService := productUsecase.NewProductService(productRepo)

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
	router.Run(":8000")
}