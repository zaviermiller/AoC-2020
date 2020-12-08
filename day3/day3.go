package day3

import (
	"strconv"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day3 struct {
}

func (d Day3) Task1() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day3/input.txt")

    return strconv.Itoa(treesFromSlope(input, 3, 1))
}

func treesFromSlope(input []string, right int, down int) int {
	lineLength := len(input[0])
	currentX := 0
	treeCounter := 0

	for i := 0; i < len(input); i += down {
		if input[i][currentX] == '#' {
			treeCounter += 1
		}

		currentX += right
		if currentX >= lineLength { currentX -= lineLength }
	}

	return treeCounter
}



func (d Day3) Task2() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day3/input.txt")

	t1 := treesFromSlope(input, 1, 1)
	t2 := treesFromSlope(input, 3, 1)
	t3 := treesFromSlope(input, 5, 1)
	t4 := treesFromSlope(input, 7, 1)
	t5 := treesFromSlope(input, 1, 2)

	return strconv.Itoa(t1 * t2 * t3 * t4 * t5)
}
