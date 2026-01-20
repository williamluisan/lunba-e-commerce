package main

import (
	"context"

	paymentMysqlRepo "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/mysql/payment"
	paymentService "github.com/williamluisan/lunba-e-commerce/internal/usecase/payment"
)

func main() {
	ctx := context.Background()

	/*
	userService := service.NewUserService(mysql.NewMysqlUserRepository())

	// dummy
	input := &entity.UserInput{
			FirstName: "john",
			LastName: "doe",
			Email: "0912student01@test.com",
			Password: "test1234",
	}

	err := userService.Create(context.Background(), input)
	if err != nil {
			fmt.Println(err)
			return
	}
	*/

	paymentService := paymentService.New(paymentMysqlRepo.New())
	orderId := int64(1234)
	paymentService.Process(ctx, orderId)
}

