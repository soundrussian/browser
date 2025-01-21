package utils

import "fmt"

func PadStringWithSpaces(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return fmt.Sprintf("%-*s", length, s)
}
