// AUTO GENERATED FILE -- DO NOT EDIT!! EDIT template_main.go AND RUN TWICE

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
	u "github.com/zaviermiller/advent-of-code-2020/util"

	"github.com/joho/godotenv"
)

type Day interface {
	Task1() string
	Task2() string
}

var loc, _ = time.LoadLocation("America/New_York")
var sessionCookie string

func main() {
	var DAYS []Day
	// AUTO GENERATED -- day_struct
	dayFlag := flag.Int("d", 0, "day of advent to run")
	// inputFlag := flag.String("inp", "", "custom input to tasks")
	flag.Parse()

	// Get env var
	err := godotenv.Load()
	check(err)

	sessionCookie = os.Getenv("AOC_COOKIE")

	fmt.Println("\nAttempting to generate new files...")
	autogen() // autogenerate new day files
	fmt.Println("Done\n")

	if *dayFlag == 0 {
		return
	}

	year, _, day := time.Now().In(loc).Date()
	if (year < 2021 || day > 25) && (*dayFlag > day || *dayFlag < 0) {
		fmt.Println("ERROR: That date is not valid (yet!)")
		return
	}
	
	dayObj := DAYS[*dayFlag - 1]

	fmt.Println("\033[0mTASK 1 ANSWER:\n\033[31;1m" + dayObj.Task1())
	fmt.Println("\033[0mTASK 2 ANSWER:\n\033[32;1m" + dayObj.Task2() + "\033[0m\n")
	
}


func autogen() {
	// Declare http client
	client := &http.Client{}
	
	cwd, _ := os.Getwd()
	year, month, day := time.Now().In(loc).Date()

	changeFlag := false
	
	for i := 1; i <= day; i++ {
		dayDir := cwd + fmt.Sprintf("/day%d",i)
		_, err := os.Open(dayDir + fmt.Sprintf("/day%d.go",i))
		if err != nil && year == 2020 && month == 12 {
			changeFlag = true
			// fmt.Println("ERROR " + err.Error())
			_, err := os.Stat(dayDir)
			if os.IsNotExist(err) {
				os.Mkdir(dayDir, 0777)
			}
			
			// ====================
			// TEXT FILE GENERATION
			// ====================
			
			// make request from AoC2020 day n
			req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2020/day/%d/input",i), nil)
			check(err)
			
			// Set cookie and do request
			req.Header.Set("Cookie", fmt.Sprintf("session=%s;",sessionCookie))
			resp, err := client.Do(req)
			check(err)
			
			// Read response
			inputData, err := ioutil.ReadAll(resp.Body)
			check(err)
			
			inputFile, err := os.Create(dayDir + "/input.txt")
			
			defer inputFile.Close()
			
			_, err = inputFile.Write(inputData)
			check(err)
			
			// ==================
			// GO FILE GENERATION
			// ==================
			
			file, err := os.Create(dayDir + fmt.Sprintf("/day%d.go",i))
			check(err)
			
			defer file.Close()
			
			_, err = file.WriteString(fmt.Sprintf("package day%d\n\nimport (\n\t// \"strconv\"\n\n\tu \"github.com/zaviermiller/advent-of-code-2020/util\"\n)\n\ntype Day%d struct {\n}\n\nfunc (d Day%d) Task1() string {\n\tinput, _ := u.InputFromFile(\"%s/input.txt\")\n\n\treturn input[0]\n}\n\nfunc (d Day%d) Task2() string {\n\tinput, _ := u.InputFromFile(\"%s/input.txt\")\n\n\treturn input[0]\n}\n", i, i, i, dayDir, i, dayDir))
			check(err)
			
			fmt.Println(fmt.Sprintf("\nGenerated day %d folder\n",i))
		}
	}

	if changeFlag {
		iMain, err := os.Open("template_main.go")
		check(err)
		
		defer iMain.Close()
		
		mainReader := bufio.NewReader(iMain)
		
		oMain, err := os.Create("main.go")
		check(err)
		
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
	}
}

func check(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error)
		os.Exit(0)
	}
}