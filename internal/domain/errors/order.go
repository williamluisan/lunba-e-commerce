package errors

import (
	"fmt"
)

type OrderAlreadyExistsError struct {
	ProductName string
}

func (e *OrderAlreadyExistsError) Error() string {
	return fmt.Sprintf(
		"order already exists for product=%s",
		e.ProductName,
	)
}