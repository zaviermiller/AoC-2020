package util

import (
	"strings"
)

func InputToSlice(input string) []string {
	arr := strings.Split(input, "\n")
	return arr[:len(arr) - 1]
}