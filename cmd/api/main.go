package main

import (
	"log"

	transHttpGin "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin"

	gormOrderRepo "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/order"
	gormMysql "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"
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

	/* transport handler */
	// userHandler := userHandler.NewUserHandler(userService)
	// productHandler := productHandler.NewProductHandler(productService)

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		// UserHandler: userHandler,
		// ProductHandler: productHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8000")
}