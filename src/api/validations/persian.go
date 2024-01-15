package validations

import (
	"github.com/go-playground/validator"
	"github.com/salmantaghooni/golang-car-web-api/src/common"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.IranianMobileNumberValidate(value)
}
