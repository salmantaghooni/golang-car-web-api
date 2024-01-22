package common

import (
	"log"
	"regexp"
)

const iranianMobileNumberPattern = `^09(1[0-9]|3[1-9])-?[0-9]{3}-?[0-9]{4}`

func IranianMobileNumberValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(iranianMobileNumberPattern, mobileNumber)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
