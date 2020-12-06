package day3
import (
	"strconv"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day3 struct {
}

func (d Day3) Task1(input string) string {
	inpArr := u.InputToSlice(input)
	lineLength := len(inpArr[0])
	currentX := 0
	treeCounter := 0

	for _, line := range inpArr {
		if line[currentX] == '#' {
			treeCounter += 1
		}

		currentX += 3
		if currentX >= lineLength { currentX -= lineLength }
	}
	
    return strconv.Itoa(treeCounter)
}

func (d Day3) Task2(input string) string {
	inpArr := u.InputToSlice(input)

	t1 := treesFromSlope(inpArr, 1, 1)
	t2 := treesFromSlope(inpArr, 3, 1)
	t3 := treesFromSlope(inpArr, 5, 1)
	t4 := treesFromSlope(inpArr, 7, 1)
	t5 := treesFromSlope(inpArr, 1, 2)

	return strconv.Itoa(t1 * t2 * t3 * t4 * t5)
}

func treesFromSlope(inpArr []string, right int, down int) int {
	lineLength := len(inpArr[0])
	currentX := 0
	treeCounter := 0

	for i := 0; i < len(inpArr); i += down {
		if inpArr[i][currentX] == '#' {
			treeCounter += 1
		}

		currentX += right
		if currentX >= lineLength { currentX -= lineLength }
	}

	return treeCounter
}


// went for speed, got sub 2000