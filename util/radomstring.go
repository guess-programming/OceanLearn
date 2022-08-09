package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM._-")
	results := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range results {
		results[i] = letters[rand.Intn(len(letters))]
	}
	return string(results)
}
