package main

import (
	rabbitMQ "github.com/williamluisan/lunba-e-commerce/internal/transport/rabbitmq"
	rabbitMQConsumer "github.com/williamluisan/lunba-e-commerce/internal/transport/rabbitmq/consumer"
)

func main() {
	// ctx := context.Background()

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

	/*
	paymentService := paymentService.New(paymentMysqlRepo.New())
	orderId := int64(1234)
	paymentService.Process(ctx, orderId)
	*/

	rabbitMQconn := rabbitMQ.Conn()
	rabbitMQConsumer.PaymentConsume(rabbitMQconn)
}

