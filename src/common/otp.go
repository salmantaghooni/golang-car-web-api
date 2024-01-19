package common

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GenereateOTP(digits int) string {
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(digits)))
	max := int(math.Pow(10, float64(digits)) - 1)

	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
