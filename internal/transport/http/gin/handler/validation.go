package handler

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) map[string][]string {
	errors := make(map[string][]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			fmt.Println(fe.Field())
			field := strings.ToLower(fe.Field())

			switch fe.Tag() {
			case "required":
				errors[field] = append(errors[field], "is required")
			case "email":
				errors[field] = append(errors[field], "must be a valid email address")
			default:
				errors[field] = append(errors[field], "is invalid")
			}
		}
	} else {
		// handle other error
		errors["body"] = append(errors["body"], err.Error())
	}

	return errors
}