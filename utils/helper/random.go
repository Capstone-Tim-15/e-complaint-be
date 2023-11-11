package helper

import (
	"math/rand"
	"time"
)

const length = 6

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())

	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	length := 6

	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func GenerateOTP() string {
	characters := "0123456789"
	otp := make([]byte, length)

	rand.Seed(time.Now().UnixNano())

	for i := range otp {
		otp[i] = characters[rand.Intn(len(characters))]
	}

	return string(otp)
}
