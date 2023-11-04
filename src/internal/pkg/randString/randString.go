package randString

import (
	"math/rand"
	"time"
)

const (
	length   = 10
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func GenerateString() string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = alphabet[r.Intn(len(alphabet))]
	}
	return string(shortKey)

}
