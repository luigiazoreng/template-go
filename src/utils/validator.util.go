package utils

import "github.com/go-playground/validator"

func Validate(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}