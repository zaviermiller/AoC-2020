package util

import (
	"strings"
)

func InputToSlice(input string) []string {
	arr := strings.Split(input, "\n")
	return arr[:len(arr) - 1]
}

func BatchInputTo2DSlice(input string) [][]string {
	arr := strings.Split(input, "\n\n")
	var finalArr [][]string

	checkFunc := func(c rune) bool {
		return c == '\n' || c == ' '
	}

	for _, item := range arr {
		subArr := strings.FieldsFunc(item, checkFunc)
		finalArr = append(finalArr, subArr)
	}

	return finalArr
}