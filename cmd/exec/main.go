package main

import (
	"context"
	"fmt"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/service"
	"github.com/williamluisan/lunba-e-commerce/internal/infrastructure/mysql"
)

func main() {
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
}

