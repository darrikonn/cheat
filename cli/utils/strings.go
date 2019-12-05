package utils

import (
	"strings"
)

// RemoveAtIndex : removes item at index from a list
func RemoveAtIndex(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// GetFirstLine : returns first line of a string
func GetFirstLine(str string) string {
	return strings.Split(str, "\n")[0]
}

// SingularOrPlural : returns the pluralization of the noun
func SingularOrPlural(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}
