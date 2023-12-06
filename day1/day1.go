package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numbersFromString []string

func main() {
	fileLines := readFile("../data/day1-1.txt")
	count := 0
	for _, line := range fileLines {
		count = count + getIntsFromString(line)
	}
	fmt.Println(count)
}

func getIntsFromString(line string) int {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	submatchall := re.FindAllString(line, -1)
	firstNum := submatchall[0]

	if len(submatchall) > 1 {
		if len(firstNum) > 1 {
			firstNum = string(getChar(firstNum, 0))
		}

		lastNum := submatchall[len(submatchall)-1]

		if len(lastNum) > 1 {
			lastNum = string(getChar(lastNum, len(lastNum)-1))
		}

		sum := firstNum + lastNum
		sumInt, _ := strconv.Atoi(sum)
		return sumInt
	}

	if len(firstNum) == 1 {
		firstNum = string(getChar(firstNum, 0))
	} else {
		firstNumInitial := string(getChar(firstNum, 0))
		lastNum := string(getChar(firstNum, len(firstNum)-1))
		sum := firstNumInitial + lastNum
		sumInt, _ := strconv.Atoi(sum)
		return sumInt
	}

	singleNum := firstNum + firstNum
	singleNumInt, _ := strconv.Atoi(singleNum)
	return singleNumInt
}

func readFile(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLine []string

	for fileScanner.Scan() {
		fileLine = append(fileLine, fileScanner.Text())
	}

	readFile.Close()

	return fileLine
}

func getChar(str string, index int) rune {
	return []rune(str)[index]
}
