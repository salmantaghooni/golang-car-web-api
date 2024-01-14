package validations

import (
	"errors"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	Property string `json:"peroperty"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var ValidationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			ValidationErrors = append(ValidationErrors, el)
		}
		return &ValidationErrors
	}
	return nil
}
