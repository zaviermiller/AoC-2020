package day4
import (
	"strings"
	"strconv"
	// "fmt"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day4 struct {
}

func (d Day4) Task1(input string) interface{} {
	inpArr := u.BatchInputTo2DSlice(input)
	validCount := 0

	for _, passport := range inpArr {
		unmetReqs := []string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		for _, item := range passport {
			temp := strings.Split(item, ":")
			desc := temp[0]
			// val := temp[1]

			unmetReqs = containsAndRemove(unmetReqs, desc)
		}
		if len(unmetReqs) == 0 {
			validCount += 1
		}

	}

    return strconv.Itoa(validCount)
}

func containsAndRemove(s []string, e string) []string {
    for i, a := range s {
        if a == e {
			s = append(s[:i],s[i+1:]...)
            return s
        }
    }
    return s
}

func checkValidity(val string, desc string) bool {
	switch desc {
	case "byr":
		if len(val) != 4 {
			return false;
		}
		
		byr,_ := strconv.Atoi(val)
		if byr < 1920 || byr > 2002 {
			return false
		}
	return true

	case "iyr":
		if len(val) != 4 {
			return false;
		}
		
		iyr,_ := strconv.Atoi(val)
		if iyr < 2010 || iyr > 2020 {
			return false
		}
	return true

	case "eyr":
		if len(val) != 4 {
			return false;
		}
		
		eyr,_ := strconv.Atoi(val)
		if eyr < 2020 || eyr > 2030 {
			return false
		}
	return true

	case "hgt":
		num, _ := strconv.Atoi(val[:len(val) - 2])
		unit := val[len(val) - 2:]

		if unit == "cm" {
			return !(num > 193 || num < 150)
		} 

		if unit == "in" {
			return !(num > 76 || num < 59)
		}

	case "hcl":
		if val[0] != '#' || len(val) != 7 {
			return false
		}

		_, err := strconv.ParseInt(val[1:], 16, 64)
		if err != nil {
			return false
		}

		return true

	case "ecl":
		colors := []string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, item := range colors {
			if item == val {
				return true
			}
		}

		return false

	case "pid":
		return len(val) == 9
	default:
		return false
	}
	return false
}

func (d Day4) Task2(input string) interface{} {
	inpArr := u.BatchInputTo2DSlice(input)
	validCount := 0

	for _, passport := range inpArr {
		unmetReqs := []string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		validityCount := 0
		for _, item := range passport {
			temp := strings.Split(item, ":")
			desc := temp[0]
			val := temp[1]

			unmetReqs = containsAndRemove(unmetReqs, desc)
			// fmt.Println(desc, val, checkValidity(val, desc))
			if checkValidity(val, desc) {
				validityCount += 1
			}
		}
		// fmt.Println(validityCount)
		if len(unmetReqs) == 0 && validityCount == 7{
			validCount += 1
		}

	}

	return strconv.Itoa(validCount)
}
