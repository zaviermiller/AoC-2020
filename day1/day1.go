package day1

import (
	"strconv"
	"fmt"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day1 struct {
}

func (d Day1) Task1() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day1/input.txt")
	intArr := make([]int, len(input))
	for i, item := range input {
		fmt.Println(item)
		intArr[i], _ = strconv.Atoi(item)
	}

	fmt.Println(intArr)

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

func (d Day1) Task2() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day1/input.txt")
	intArr := make([]int, len(input))
	for i, item := range input {
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