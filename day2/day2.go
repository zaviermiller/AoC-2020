package day2

import (
	"strconv"
	"strings"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day2 struct {
}

func (d Day2) Task1() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day2/input.txt")
	counter := 0
    for _, item := range input {
        stringArr := strings.SplitN(item, "-", 2)
        firstVal, _ := strconv.Atoi(stringArr[0])
        theRest := stringArr[1]
        stringArr = strings.SplitN(theRest, " ", 2)
        secondVal, _ := strconv.Atoi(stringArr[0])
        theRest = stringArr[1]
        stringArr = strings.SplitN(theRest, ":", 2)
        searchChar := []rune(stringArr[0])[0]
        theRest = stringArr[1]
        password := strings.Split(theRest, " ")[1]

        if checkValidity((firstVal), (secondVal), (searchChar), password) {
            counter += 1
        }
    }

    return strconv.Itoa(counter)
}

func checkValidity(num1 int, num2 int, search rune, password string) bool {
    counter := 0
    for _, c := range password {
        if c == search {
            counter += 1
        }
    }
    // fmt.Println(num1, num2, string(search), password, counter)

    if counter >= num1 && counter <= num2 {
        return true
    }
    return false
}

func (d Day2) Task2() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day2/input.txt")
	counter := 0

    for _, item := range input {
        stringArr := strings.SplitN(item, "-", 2)
        firstVal, _ := strconv.Atoi(stringArr[0])
        theRest := stringArr[1]
        stringArr = strings.SplitN(theRest, " ", 2)
        secondVal, _ := strconv.Atoi(stringArr[0])
        theRest = stringArr[1]
        stringArr = strings.SplitN(theRest, ":", 2)
        searchChar := []byte(stringArr[0])[0]
        theRest = stringArr[1]
        password := strings.Split(theRest, " ")[1]

        if checkValidity2(firstVal, secondVal, searchChar, password) {
            counter += 1
        }
    }
	return strconv.Itoa(counter)
}

func checkValidity2(num1 int, num2 int, search byte, password string) bool {
    return (password[num1 - 1] == search) != (password[num2 - 1] == search)
}