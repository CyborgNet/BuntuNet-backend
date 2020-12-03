package utils

import (
	"math/rand"
	"time"
)

func RandomString(strlen int, icint bool) string { //Generates a random string
	if icint != false {
		rand.Seed(time.Now().UTC().UnixNano())
		const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
		result := make([]byte, strlen)
		for i := 0; i < strlen; i++ {
			result[i] = chars[rand.Intn(len(chars))]
		}
		return string(result)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
