package main

import (
	"log"

	transHttpGin "lunba-e-commerce/internal/transport/http/gin"

	gormMysql "lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"
	gormOrderRepo "lunba-e-commerce/internal/infrastructure/gorm/repository/order"

	orderUseCase "lunba-e-commerce/internal/usecase/order"

	orderHandler "lunba-e-commerce/internal/transport/http/gin/handler/order"
)

func main() {
	dsn := "root:toor@tcp(mysql:3306)/ln-e-commerce_order_payment?parseTime=true"
	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	/* infrastructure */
	// mysqlUser := mysql.NewMysqlUserRepository()
	// gormUserRepo := gormUserRepo.NewUserRepository(db)
	// gormProductRepo := gormProductRepo.NewProductRepository(db)
	gormOrderRepo := gormOrderRepo.New(db)

	/* domain */
	// userRepo := gormUserRepo
	// userService := userUsecase.NewUserService(userRepo)
	// productRepo := gormProductRepo
	// productService := productUsecase.NewProductService(productRepo)
	orderRepo := gormOrderRepo
	orderService := orderUseCase.New(orderRepo)

	/* transport handler */
	// userHandler := userHandler.NewUserHandler(userService)
	// productHandler := productHandler.NewProductHandler(productService)
	orderHandler := orderHandler.New(orderService)

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		// UserHandler: userHandler,
		// ProductHandler: productHandler,
		OrderHandler: orderHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8000")
}