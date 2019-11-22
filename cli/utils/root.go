package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomID() string {
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Intn(1000000)) // 10 ^ 6
	return strings.Repeat("0", 6-len(id)) + id
}

func SingularOrPlural(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
