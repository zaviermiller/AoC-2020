package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"flag"
	"time"

	"github.com/zaviermiller/advent-of-code-2020/day1"
	"github.com/zaviermiller/advent-of-code-2020/day2"
	"github.com/zaviermiller/advent-of-code-2020/day3"

	"github.com/joho/godotenv"
)

type Day interface {
	Task1(string) interface{}
	Task2(string) interface{}
}

func main() {
	var DAYS []Day
	DAYS = append(DAYS, &day1.Day1{}, &day2.Day2{}, &day3.Day3{})
	dayFlag := flag.Int("d", 0, "day of advent to run")
	flag.Parse()

	fmt.Println("\nAttempting to generate new files...")
	autogen() // autogenerate new day files
	fmt.Println("Done\n")

	year, _, day := time.Now().Date()
	if (year < 2021 || day > 25) && (*dayFlag > day || *dayFlag < 1) {
		fmt.Println("ERROR: That date is not valid (yet!)")
		return
	}

	input := getInputOnDay(*dayFlag)
	
	dayObj := DAYS[*dayFlag - 1]

	fmt.Println("TASK 1 ANSWER:\n\033[1m" + dayObj.Task1(input).(string))
	fmt.Println("\033[0mTASK 2 ANSWER:\n\033[1m" + dayObj.Task2(input).(string) + "\033[0m\n")
	
}

func autogen() {
	cwd, _ := os.Getwd()
	year, month, day := time.Now().Date()

	for i := 1; i <= day; i++ {
		dayDir := cwd + fmt.Sprintf("/day%d",i)
		_, err := os.Open(dayDir + fmt.Sprintf("/day%d.go",i))
		if err != nil && year == 2020 && month == 12 {
			// fmt.Println("ERROR " + err.Error())
			_, err := os.Stat(dayDir)
			if os.IsNotExist(err) {
				os.Mkdir(dayDir, 0777)
			}

			file, err := os.Create(dayDir + fmt.Sprintf("/day%d.go",i))
			check(err)

			defer file.Close()

			_, err = file.WriteString(fmt.Sprintf("package day%d\nimport (\n\n)\n\ntype Day%d struct {\n}\n\nfunc (d Day%d) Task1(input string) interface{} {\n    return input\n}\n\nfunc (d Day%d) Task2(input string) interface{} {\n	return input\n}\n", i, i, i, i))
			check(err)

			fmt.Println(fmt.Sprintf("\nGenerated day%d\n",i))
		}
	}
}

func getInputOnDay(day int) (string) {
	// Declare http client
	client := &http.Client{}
	
	// Declare HTTP Method and Url
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2020/day/%d/input",day), nil)
	check(err)

	// Get env var
	err = godotenv.Load()
	check(err)

	sessionCookie := os.Getenv("AOC_COOKIE")

	// Set cookie
	req.Header.Set("Cookie", fmt.Sprintf("session=%s;",sessionCookie))
	resp, err := client.Do(req)
	check(err)
	// Read response
	data, err := ioutil.ReadAll(resp.Body)
	check(err)

	// Return response
	return string(data)
}

func check(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error)
		os.Exit(0)
	}
}