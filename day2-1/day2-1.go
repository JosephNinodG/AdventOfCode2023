package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxRedCubes int = 12
const maxGreenCubes int = 13
const maxBlueCubes int = 14

func main() {
	fileLines := readFile("../data/day2.txt")
	count := 0
	for _, line := range fileLines {
		newNum := getIntsFromString(line)
		count = count + newNum
	}
	fmt.Println(count)
}

func getIntsFromString(line string) int {
	splitSliceID := strings.Split(line, ":")
	gameID := getGameID(splitSliceID)
	isPossible := getCubeNumbers(strings.Split(splitSliceID[1], ";"))

	if isPossible {
		return gameID
	}

	return 0
}

func getCubeNumbers(game []string) bool {
	for _, round := range game {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			cubeInfo := strings.Split(cube, " ")
			cubeNum := 0
			cubeColour := ""

			if cubeInfo[0] == "" {
				cubeNum, _ = strconv.Atoi(cubeInfo[1])
				cubeColour = cubeInfo[2]
			} else {
				cubeNum, _ = strconv.Atoi(cubeInfo[0])
				cubeColour = cubeInfo[1]
			}

			switch cubeColour {
			case "red":
				if cubeNum > maxRedCubes {
					return false
				}
			case "green":
				if cubeNum > maxGreenCubes {
					return false
				}
			case "blue":
				if cubeNum > maxBlueCubes {
					return false
				}
			}
		}
	}

	return true
}

func getGameID(splitSlice []string) int {
	gameIDSlice := strings.Split(splitSlice[0], " ")
	gameID, _ := strconv.Atoi(gameIDSlice[1])
	return gameID
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
