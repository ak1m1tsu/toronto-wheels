package validator

import (
	v "github.com/go-playground/validator/v10"
)

var (
	validator = v.New()
)

func Validate(data interface{}) error {
	return validator.Struct(data)
}
