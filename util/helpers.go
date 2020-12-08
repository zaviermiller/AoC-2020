package util

import (
	"strings"
	"os"
	"io"
	// "fmt"
	"bufio"
)

var SearchString string = "// AUTO GENERATED --"

func InputFromFile(filePath string) ([]string, error) {
	output := []string{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	mainReader := bufio.NewReader(file)

	for {
		line, err := mainReader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break;
		}

		
		output = append(output, line[:len(line) - 1])
	}
	
	return output, nil
}

func InputStringFromFile(filePath string) (string, error) {
	output := ""
	file, err := os.Open(filePath)
	if (err != nil) {
		return "", err
	}

	defer file.Close()

	mainReader := bufio.NewReader(file)

	for {
		line, err := mainReader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break;
		}

		output += line
	}

	return output, nil
}

func InputToSlice(input string) []string {
	arr := strings.Split(input, "\n")
	return arr[:len(arr) - 1]
}

func BatchInputTo2DSlice(input string) [][]string {
	arr := strings.Split(input, "\n\n")
	var finalArr [][]string

	checkFunc := func(c rune) bool {
		return c == '\n' || c == ' '
	}

	for _, item := range arr {
		subArr := strings.FieldsFunc(item, checkFunc)
		finalArr = append(finalArr, subArr)
	}

	return finalArr
}