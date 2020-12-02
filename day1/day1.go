package day1

import (
	// "fmt"
	"strconv"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day1 struct {
}

func (d Day1) Task1(input string) interface{} {
	inpArr := u.InputToSlice(input)
	intArr := make([]int, len(inpArr))
	for i, item := range inpArr {
		intArr[i], _ = strconv.Atoi(item)
	}

	solution := strconv.Itoa(findSum2020(intArr))
	
	return solution
}

// brute force B) 100 emoji
func findSum2020(arr []int) int {
	for _, num1 := range arr {
		for _, num2 := range arr {
			if num1 + num2 == 2020 {
				return num1 * num2
			}
		}
	}
	return -1
}

// woah i can even brute force this one :o
func (d Day1) Task2(input string) interface{} {
	inpArr := u.InputToSlice(input)
	intArr := make([]int, len(inpArr))
	for i, item := range inpArr {
		intArr[i], _ = strconv.Atoi(item)
	}

	solution := strconv.Itoa(findTripleSum2020(intArr))
	
	return solution
}

func findTripleSum2020(arr []int) int {
	for _, num1 := range arr {
		for _, num2 := range arr {
			if num1 + num2 < 2020 {
				for _, num3 := range arr {
					if num1 + num2 + num3 == 2020 {
						return num1 * num2 * num3
					}
				}
			}
		}
	}
	return -1
}