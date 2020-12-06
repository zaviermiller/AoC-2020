package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bufio"
	"os"
	"flag"
	"strings"
	"time"
	"io"

	// AUTO GENERATED -- imports
	"github.com/zaviermiller/advent-of-code-2020/day1"
	"github.com/zaviermiller/advent-of-code-2020/day2"
	"github.com/zaviermiller/advent-of-code-2020/day3"
	"github.com/zaviermiller/advent-of-code-2020/day4"
	"github.com/zaviermiller/advent-of-code-2020/day5"
	u "github.com/zaviermiller/advent-of-code-2020/util"

	"github.com/joho/godotenv"
)

type Day interface {
	Task1(string) string
	Task2(string) string
}

var loc, _ = time.LoadLocation("America/New_York")

func main() {
	var DAYS []Day
	// AUTO GENERATED -- day_struct
	DAYS = append(DAYS, &day1.Day1{}, &day2.Day2{}, &day3.Day3{}, &day4.Day4{}, &day5.Day5{})
	dayFlag := flag.Int("d", 0, "day of advent to run")
	inputFlag := flag.String("inp", "", "custom input to tasks")
	flag.Parse()

	fmt.Println("\nAttempting to generate new files...")
	autogen() // autogenerate new day files
	fmt.Println("Done\n")

	year, _, day := time.Now().In(loc).Date()
	if (year < 2021 || day > 25) && (*dayFlag > day || *dayFlag < 1) {
		fmt.Println("ERROR: That date is not valid (yet!)")
		return
	}

	var input string

	if len(*inputFlag) == 0 {
		input = getInputOnDay(*dayFlag)
	} else {
		input = *inputFlag
	}
	
	dayObj := DAYS[*dayFlag - 1]

	fmt.Println("TASK 1 ANSWER:\n\033[1m" + dayObj.Task1(input))
	fmt.Println("\033[0mTASK 2 ANSWER:\n\033[1m" + dayObj.Task2(input) + "\033[0m\n")
	
}

func autogen() {
	cwd, _ := os.Getwd()
	year, month, day := time.Now().In(loc).Date()

	iMain, err := os.Open("main.template.go")
	check(err)

	defer iMain.Close()

	mainReader := bufio.NewReader(iMain)

	oMain, _ := os.Create("main.go")

	defer oMain.Close()

	for {
		line, err := mainReader.ReadString('\n')

								// LMFAOOOOOOO cant have real search string in file
		if strings.Contains(line, u.SearchString) {
			strippedLine := strings.Replace(strings.TrimSpace(line), " ", "", -1)
			part := strings.Split(strippedLine, "--")[1]

			switch part {
			case "imports":
				for i := 1; i <= day; i++ {
					line += fmt.Sprintf("\t\"github.com/zaviermiller/advent-of-code-2020/day%d\"\n",i)
				}
			case "day_struct":
				line += "\tDAYS = append(DAYS, "
				for i := 1; i <= day; i++ {
					line += fmt.Sprintf("&day%d.Day%d{}",i,i)
					if i != day {
						line += ", "
					}
				}
				line += ")\n"
			}

		}

        oMain.WriteString(line)
        if err != nil {
            if err != io.EOF {
                fmt.Println("error:", err)
            }
            break
        }
    }

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

			_, err = file.WriteString(fmt.Sprintf("package day%d\n\nimport (\n\t// strconv\n\n\tu \"github.com/zaviermiller/advent-of-code-2020/util\"\n)\n\ntype Day%d struct {\n}\n\nfunc (d Day%d) Task1(input string) string {\n\treturn input\n}\n\nfunc (d Day%d) Task2(input string) string {\n	return input\n}\n", i, i, i, i))
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