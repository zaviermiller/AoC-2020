package day7

import (
	"strconv"
	"strings"

	u "github.com/zaviermiller/advent-of-code-2020/util"
)

type Day7 struct {
}

type Bag struct {
	Desc  	 string
	Children map[string]int
}

func (d Day7) Task1() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day7/input.txt")
	// parse input
	bags := parseInput(input)

	knownContains := []string {}
	goodBags := 0

	for _, bag := range(bags) {
		if recursivelyTestBag(bag, bags, &knownContains) {
			goodBags += 1
			// fmt.Println(bag.Desc, knownContains)
		}
	}
	return strconv.Itoa(goodBags)
}

func recursivelyTestBag(bag *Bag, bags map[string]*Bag, knowns *[]string) bool {
	if bag == nil {
		return false
	}
	for desc, _ := range bag.Children {
		if desc == "shiny gold" || contains(*knowns, desc) || recursivelyTestBag(bags[desc], bags, knowns) {
			if !contains(*knowns, bag.Desc) { *knowns = append(*knowns, bag.Desc) }
			return true
		}
	}

	return false
} 

func contains(knowns []string, target string) bool {
	for _, item := range knowns {
		if item == target {
			return true
		}
	}
	return false
}

func parseInput(inpArr []string) map[string]*Bag {
	bags := map[string]*Bag{}

	for _, item := range inpArr {
		splitItem := strings.Split(item, " bags contain ")
		bag := &Bag { Desc: splitItem[0] }
		innerBagString := splitItem[1]
		if splitItem[1] == "no other bags." {
			bags[bag.Desc] = bag
			continue
		}
		innerBags := map[string]int{}
		for {
			splitItem = strings.SplitN(innerBagString, ", ", 2)
			if innerBagString == splitItem[0] {
				lastItem := strings.Split(innerBagString,".")[0]
				count, _ := strconv.Atoi(string(lastItem[0]))
				var desc string
				if count == 1 {
					desc = strings.Split(lastItem[2:], " bag")[0]
				} else {
					desc = strings.Split(lastItem[2:], " bags")[0]
				}
				innerBags[desc] = count
				break;
			}
			count, _ := strconv.Atoi(string(splitItem[0][0]))
			var desc string
			if count == 1 {
				desc = strings.Split(splitItem[0][2:], " bag")[0]
			} else {
				desc = strings.Split(splitItem[0][2:], " bags")[0]
			}
			innerBagString = splitItem[1]
			innerBags[desc] = count
		}
		bag.Children = innerBags
		bags[bag.Desc] = bag

	}

	return bags
}

func (d Day7) Task2() string {
	input, _ := u.InputFromFile("/home/zavier/go/src/github.com/zaviermiller/advent-of-code-2020/day7/input.txt")

	bags := parseInput(input)
	
	return strconv.Itoa(recursiveCountChildren(bags["shiny gold"], bags))
}

func recursiveCountChildren(bag *Bag, bags map[string]*Bag) int {
	if bag == nil {
		return 0
	}

	count := 0

	for desc, c := range bag.Children {
		count += c + c * recursiveCountChildren(bags[desc], bags)
	}

	return count
} 