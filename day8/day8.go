package day8

import (
	"strconv"
	"strings"
	"fmt"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day8 struct {
}

func (d Day8) Task1() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day8/input.txt")

	accumulator := 0
	currentInstruction := 0
	seen := []int {}

	for {
		if contains(seen, currentInstruction) {
			break
		}
		item := input[currentInstruction]
		split := strings.Split(item, " ")
		val, _ := strconv.Atoi(split[1])
		op := split[0]
		runOp(op, val, &accumulator, &currentInstruction, &seen)
	}

	return strconv.Itoa(accumulator)
}

func runOp(op string, val int, acc *int, inst *int, seen *[]int) {
	*seen = append(*seen, *inst)
	switch op {
	case "acc":
		*acc += val
		*inst += 1
		break
	case "jmp":
		*inst += val
		break
	default:
		*inst += 1
		break
	}
}

func contains(arr []int, target int) bool {
	for _, item := range arr {
		if target == item {
			return true
		}
	}
	return false
}

// hmmm
func (d Day8) Task2() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day8/input.txt")

	accumulator := 0
	currentInstruction := 0
	seen := []int {}

	for {
		if contains(seen, currentInstruction) {
			break
		}
		item := input[currentInstruction]
		split := strings.Split(item, " ")
		val, _ := strconv.Atoi(split[1])
		op := split[0]
		runOp(op, val, &accumulator, &currentInstruction, &seen)
	}

	fmt.Println(seen)

	for i := len(seen) - 1; i >= 0; i-- {
		inputCopy := make([]string, len(input))
		copy(inputCopy, input)

		currentChange := inputCopy[seen[i]]
		split := strings.Split(currentChange, " ")
		val := split[1]
		op := split[0]
		if op == "nop" {
			inputCopy[seen[i]] = "jmp " + val
		} else if op == "jmp" {
			inputCopy[seen[i]] = "nop " + val
		}

		isInfinite, a := testInfinite(inputCopy) 
		if !isInfinite{
			return strconv.Itoa(a)
		}
	}

	return strconv.Itoa(-1)
}

func testInfinite(input []string) (bool, int) {
	accumulator := 0
	currentInstruction := 0
	seen := []int {}

	for {
		if contains(seen, currentInstruction) {
			return true, 0
		}
		if currentInstruction >= len(input) {
			break
		}
		item := input[currentInstruction]
		split := strings.Split(item, " ")
		val, _ := strconv.Atoi(split[1])
		op := split[0]
		runOp(op, val, &accumulator, &currentInstruction, &seen)
	}
	return false, accumulator
}