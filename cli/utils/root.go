package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateRandomID : randomly generates a random 6 digit integer, casted to string
func GenerateRandomID() string {
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Intn(1000000)) // 10 ^ 6
	return strings.Repeat("0", 6-len(id)) + id
}

// SingularOrPlural : returns the pluralization of the noun
func SingularOrPlural(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}

// Check : handle errors for defer functions
func Check(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
