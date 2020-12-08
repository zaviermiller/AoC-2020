package day6

import (
	"strconv"
	"strings"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day6 struct {
}

func (d Day6) Task1() string {
	input, _ := u.InputStringFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day6/input.txt")
	groups := strings.Split(input, "\n\n")
	sumTotals := 0
	// groups := u.InputToSlice(input)
	for _, group := range groups {
		charSet := make(map[rune]bool)
		for _, char := range group {
			if char != '\n' {
				charSet[char] = true
			}
		}
		sumTotals += len(charSet)
	}
	return strconv.Itoa(sumTotals)
}

func (d Day6) Task2() string {
	input, _ := u.InputStringFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day6/input.txt")
	groups := strings.Split(input, "\n\n")
	sumTotals := 0
	// groups := u.InputToSlice(input)
	for _, group := range groups {
		numInGroup := removeEmpties(strings.Split(group, "\n"))
		charCounter := make(map[rune]int)
		for _, char := range group {
			if char != '\n' {
				charCounter[char] += 1
			}
		}

		// fmt.Println(charCounter, numInGroup, group)
		for _, val := range charCounter {
			if val == len(numInGroup) {
				sumTotals += 1
			}
		}
	}
	return strconv.Itoa(sumTotals)
}

// my answer was slightly off due to empty strings in the length array.
// and this was the quickest / easiest fix
func removeEmpties(arr []string) []string {
	r := []string{}
	for _, item := range arr {
		if item != "" {
			r = append(r, item)
		}
	}

	return r
}