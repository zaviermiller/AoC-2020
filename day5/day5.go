package day5

import (
	"strconv"
	"math"
	// "fmt"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day5 struct {
}

func (d Day5) Task1(input string) string {
	inpArr := u.InputToSlice(input)
	highestId := 0

	for _, item := range inpArr {
		row := 0
		col := 0

		for j, letter := range item {
			if j < 7 && letter == 'B' {
				row += intPower(2, 6 - j)
			} else if j >= 7 && letter == 'R' {
				col += intPower(2, 2 - (j - 7))
			}
			id := row * 8 + col
			if id > highestId {
				highestId = id
			}
		}
	}
	return strconv.Itoa(highestId)
}

func intPower(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func (d Day5) Task2(input string) string {
	inpArr := u.InputToSlice(input)
	var strs [128][8]string

	for _, item := range inpArr {
		row := 0
		col := 0

		for j, letter := range item {
			if j < 7 && letter == 'B' {
				row += intPower(2, 6 - j)
			} else if j >= 7 && letter == 'R' {
				col += intPower(2, 2 - (j - 7))
			}
			strs[row][col] = item
		}
	}

	for row, rowContent := range strs {
		for col, item := range rowContent {
			if len(item) < 10 {
				// fmt.Println(row, col ,"<<<<<<<<<")
				if (row > 10 && row < 100) {
					return strconv.Itoa(row * 8 + col)
				}
			}
		}
	}
	return "-1"
}
