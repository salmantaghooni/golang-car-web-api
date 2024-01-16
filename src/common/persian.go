package common

import (
	"regexp"

	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

const IranianMobileNumberPattern = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9]) [0-9]{7}$`

func IranianMobileNumberValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(IranianMobileNumberPattern, mobileNumber)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
