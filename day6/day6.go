package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./day6/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentMap := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		values := []string{}
		for i := 0; i < len(line); i++ {
			values = append(values, string(line[i]))
		}
		currentMap = append(currentMap, values)
	}

	res1 := problem1(copyMap(currentMap))
	fmt.Println(res1)

	res2 := problem2(copyMap(currentMap))
	fmt.Println(res2)
}

func problem1(currentMap [][]string) int {
	currRow := 0
	currColumn := 0
	totalDiff := 0
	found := false
	for i := 0; i < len(currentMap); i++ {
		for j := 0; j < len(currentMap[i]); j++ {
			if currentMap[i][j] == "^" {
				currRow = i
				currColumn = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	dir := "up"
	for currRow >= 0 && currColumn >= 0 && currRow < len(currentMap) && currColumn < len(currentMap[0]) {
		//for row := 0; row < len(currentMap); row++ {
		//	for column := 0; column < len(currentMap[0]); column++ {
		//		fmt.Print(currentMap[row][column], " ")
		//	}
		//	fmt.Print("\n")
		//}
		//fmt.Print("\n\n")
		if currentMap[currRow][currColumn] != "X" && currentMap[currRow][currColumn] != "#" {
			totalDiff++
		}

		nextRow := 0
		nextCol := 0
		switch dir {
		case "up":
			nextRow = currRow - 1
			nextCol = currColumn
		case "down":
			nextRow = currRow + 1
			nextCol = currColumn
		case "left":
			nextRow = currRow
			nextCol = currColumn - 1
		case "right":
			nextRow = currRow
			nextCol = currColumn + 1
		}

		if nextRow < 0 || nextCol < 0 || nextRow >= len(currentMap) || nextCol >= len(currentMap[0]) {
			return totalDiff
		} else {
			currentMap[currRow][currColumn] = "X"
		}

		if currentMap[nextRow][nextCol] == "#" {
			switch dir {
			case "up":
				dir = "right"
				nextRow = currRow
				nextCol = currColumn + 1
			case "down":
				dir = "left"
				nextRow = currRow
				nextCol = currColumn - 1
			case "left":
				dir = "up"
				nextRow = currRow - 1
				nextCol = currColumn
			case "right":
				dir = "down"
				nextRow = currRow + 1
				nextCol = currColumn
			}
		}
		currRow = nextRow
		currColumn = nextCol
	}
	return totalDiff
}

func problem2(currentMap [][]string) int {
	currRow := 0
	currColumn := 0

	found := false
	for i := 0; i < len(currentMap); i++ {
		for j := 0; j < len(currentMap[i]); j++ {
			if currentMap[i][j] == "^" {
				currRow = i
				currColumn = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	totalLoops := 0

	for i := 0; i < len(currentMap); i++ {
		for j := 0; j < len(currentMap[i]); j++ {
			dup := copyMap(currentMap)
			if currentMap[i][j] == "." {
				dup[i][j] = "#"
			} else {
				continue
			}
			totalLoops += findLoops(dup, currRow, currColumn)
		}
	}
	return totalLoops
}

func findLoops(currentMap [][]string, currRow int, currColumn int) int {
	dir := "up"
	for currRow >= 0 && currColumn >= 0 && currRow < len(currentMap) && currColumn < len(currentMap[0]) {
		nextRow := 0
		nextCol := 0

		switch dir {
		case "up":
			nextRow = currRow - 1
			nextCol = currColumn
		case "down":
			nextRow = currRow + 1
			nextCol = currColumn
		case "left":
			nextRow = currRow
			nextCol = currColumn - 1
		case "right":
			nextRow = currRow
			nextCol = currColumn + 1
		}
		if nextRow < 0 || nextCol < 0 || nextRow >= len(currentMap) || nextCol >= len(currentMap[0]) {
			return 0
		}

		if currentMap[currRow][currColumn] == "." || currentMap[currRow][currColumn] == "^" {
			currentMap[currRow][currColumn] = "1"
		} else {
			numPass, _ := strconv.Atoi(currentMap[currRow][currColumn])
			if numPass > 4 {
				return 1
			}
			currentMap[currRow][currColumn] = strconv.Itoa(numPass + 1)
		}

		for currentMap[nextRow][nextCol] == "#" {
			switch dir {
			case "up":
				dir = "right"
				nextRow = currRow
				nextCol = currColumn + 1
			case "down":
				dir = "left"
				nextRow = currRow
				nextCol = currColumn - 1
			case "left":
				dir = "up"
				nextRow = currRow - 1
				nextCol = currColumn
			case "right":
				dir = "down"
				nextRow = currRow + 1
				nextCol = currColumn
			}
			if nextRow < 0 || nextCol < 0 || nextRow >= len(currentMap) || nextCol >= len(currentMap[0]) {
				return 0
			}
		}
		currRow = nextRow
		currColumn = nextCol

	}
	return 0
}

func copyMap(currentMap [][]string) [][]string {
	duplicate := make([][]string, len(currentMap))
	for i := range currentMap {
		duplicate[i] = make([]string, len(currentMap[i]))
		copy(duplicate[i], currentMap[i])
	}
	return duplicate
}
