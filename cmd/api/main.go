package main

import (
	"fmt"
	"log"

	config "lunba-e-commerce/config"

	transHttpGin "lunba-e-commerce/internal/transport/http/gin"

	gormMysql "lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"
	gormOrderRepo "lunba-e-commerce/internal/infrastructure/gorm/repository/order"

	orderUseCase "lunba-e-commerce/internal/usecase/order"

	orderHandler "lunba-e-commerce/internal/transport/http/gin/handler/order"

	internalUserProductService "lunba-e-commerce/internal/infrastructure/httpclient/userproduct"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var dsn string
	if (cfg.DB.Driver == "mysql") {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
			cfg.DB.User, cfg.DB.Pass, cfg.DB.Driver, cfg.DB.Port, cfg.DB.Name, cfg.DB.Charset,
		)
	}
	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	/* infrastructure */
	// mysqlUser := mysql.NewMysqlUserRepository()
	// gormUserRepo := gormUserRepo.NewUserRepository(db)
	// gormProductRepo := gormProductRepo.NewProductRepository(db)
	gormOrderRepo := gormOrderRepo.New(db)
	userProductServiceUser := internalUserProductService.NewUser()

	/* domain */
	// userRepo := gormUserRepo
	// userService := userUsecase.NewUserService(userRepo)
	// productRepo := gormProductRepo
	// productService := productUsecase.NewProductService(productRepo)
	orderRepo := gormOrderRepo
	orderService := orderUseCase.New(orderRepo, userProductServiceUser)

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