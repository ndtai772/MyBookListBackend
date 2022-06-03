package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	ALPHABET = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(ALPHABET)

	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
