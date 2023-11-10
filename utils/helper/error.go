package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ValidationError(ctx echo.Context, err error) error {
	validationError, ok := err.(validator.ValidationErrors)
	if ok {
		messages := make([]string, 0)
		for _, e := range validationError {
			messages = append(messages, fmt.Sprintf("validation error on field %s, tag %s", e.Field(), e.Tag()))
		}

		return fmt.Errorf("validation failed: %s", strings.Join(messages, "; "))
	}

	return nil
}
