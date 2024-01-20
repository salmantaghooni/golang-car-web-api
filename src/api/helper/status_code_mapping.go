package helper

import (
	"net/http"

	"github.com/salmantaghooni/golang-car-web-api/src/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	service_errors.OTPExists:   http.StatusConflict,
	service_errors.OTPUsed:     http.StatusConflict,
	service_errors.OTPNotValid: http.StatusBadRequest,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
