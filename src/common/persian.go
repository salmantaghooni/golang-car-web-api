package common

import (
	"regexp"
)

const IranianMobileNumberPattern = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9]) [0-9]{7}$`

func IranianMobileNumberValidate(mobileNumber string) bool {
	res, _ := regexp.MatchString(IranianMobileNumberPattern, mobileNumber)
	return res
}
