package handler

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())

			switch fe.Tag() {
			case "required":
				errors[field] = "is required"
			case "email":
				errors[field] = "must be a valid email address"
			default:
				errors[field] = "is invalid"
			}
		}
	}

	return errors
}