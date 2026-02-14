package messages

import (
	"fmt"
)

var (
	OrderCreated = ErrorCodeMsg{Code: "ORDER_CREATED", Message: "Order created successfully."}
)

type OrderAlreadyExistsError struct {
	ProductName string
}

func (e *OrderAlreadyExistsError) Error() string {
	return fmt.Sprintf(
		"Order already exists for product=%s",
		e.ProductName,
	)
}