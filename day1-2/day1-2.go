package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var numbersFromString []string

func main() {
	fileLines := readFile("../data/day1.txt")
	count := 0
	for _, line := range fileLines {
		newNum := getIntsFromString(line)
		count = count + newNum
	}
	fmt.Println(count)
}

func getIntsFromString(line string) int {
	numbersAsInt := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numbersAsString := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9", "zero": "0"}

	countIntPos := []int{}
	countStringPos := make(map[int]string)

	firstNum := ""
	lastNum := ""
	sumInt := 0

	for i := 0; i < len(line); i++ {
		for _, number := range numbersAsInt {
			if string(line[i]) == number {
				countIntPos = append(countIntPos, i)
			}
		}
	}

	for j := 0; j < len(line); j++ {
		for key, number := range numbersAsString {

			if (len(line) - j) >= 5 {
				stringCheck := line[j : j+5]
				if stringCheck == key {
					countStringPos[j] = number
				}
			}

			if (len(line) - j) >= 4 {
				stringCheck := line[j : j+4]
				if stringCheck == key {
					countStringPos[j] = number
				}
			}

			if (len(line) - j) >= 3 {
				stringCheck := line[j : j+3]
				if stringCheck == key {
					countStringPos[j] = number
				}
			}
		}
	}

	keys := make([]int, 0, len(countStringPos))
	for pos := range countStringPos {
		keys = append(keys, pos)
	}

	sort.Ints(countIntPos)
	sort.Ints(keys)

	if len(countStringPos) > 0 && len(countIntPos) > 0 {
		if keys[0] < countIntPos[0] {

			firstNum = countStringPos[keys[0]]
		} else {
			firstNum = string(line[countIntPos[0]])
		}

		if keys[len(keys)-1] > countIntPos[len(countIntPos)-1] {
			lastNum = countStringPos[keys[len(keys)-1]]
		} else {
			lastNum = string(line[countIntPos[len(countIntPos)-1]])
		}

		sumInt, _ = strconv.Atoi(firstNum + lastNum)

	} else if len(countStringPos) == 0 {
		if len(countIntPos) == 1 {
			sumInt, _ = strconv.Atoi(string(line[countIntPos[0]]) + string(line[countIntPos[0]]))
			return sumInt
		}

		firstNum = string(line[countIntPos[0]])
		lastNum = string(line[countIntPos[len(countIntPos)-1]])

		sumInt, _ = strconv.Atoi(firstNum + lastNum)
	} else if len(countIntPos) == 0 {
		if len(countStringPos) == 1 {
			sumInt, _ = strconv.Atoi(countStringPos[keys[0]] + countStringPos[keys[0]])
			return sumInt
		}

		firstNum = countStringPos[keys[0]]
		lastNum = countStringPos[keys[len(keys)-1]]

		sumInt, _ = strconv.Atoi(firstNum + lastNum)
	}

	return sumInt

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
