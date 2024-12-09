package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day4/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	problemMatrix := make([][]string, 140)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		problemMatrix[lineNum] = make([]string, len(line))

		for j := 0; j < len(line); j++ {
			problemMatrix[lineNum][j] = string(line[j])
		}
		lineNum++
	}

	problem1 := solveProblem1(problemMatrix)
	fmt.Println(problem1)

	problem2 := solveProblem2(problemMatrix)
	fmt.Println(problem2)
}

func solveProblem1(problemMatrix [][]string) int {
	count := 0
	for i := 0; i < len(problemMatrix); i++ {
		for j := 0; j < len(problemMatrix[i]); j++ {
			if problemMatrix[i][j] == "X" {
				count += checkNextStr(problemMatrix, i+1, j+1, "M", "++") +
					checkNextStr(problemMatrix, i+1, j, "M", "+=") +
					checkNextStr(problemMatrix, i, j+1, "M", "=+") +
					checkNextStr(problemMatrix, i-1, j-1, "M", "--") +
					checkNextStr(problemMatrix, i-1, j, "M", "-=") +
					checkNextStr(problemMatrix, i, j-1, "M", "=-") +
					checkNextStr(problemMatrix, i+1, j-1, "M", "+-") +
					checkNextStr(problemMatrix, i-1, j+1, "M", "-+")

			}
		}
	}
	return count
}

func checkNextStr(matrix [][]string, row int, col int, nextStr string, dir string) int {
	if row >= len(matrix) || row < 0 || col >= len(matrix[row]) || col < 0 {
		return 0
	}
	if matrix[row][col] == nextStr {
		if nextStr == "S" {
			return 1
		} else {
			switch nextStr {
			case "M":
				nextStr = "A"
			case "A":
				nextStr = "S"
			}
			switch dir {
			case "++":
				return checkNextStr(matrix, row+1, col+1, nextStr, dir)
			case "+=":
				return checkNextStr(matrix, row+1, col, nextStr, dir)
			case "=+":
				return checkNextStr(matrix, row, col+1, nextStr, dir)
			case "--":
				return checkNextStr(matrix, row-1, col-1, nextStr, dir)
			case "-=":
				return checkNextStr(matrix, row-1, col, nextStr, dir)
			case "=-":
				return checkNextStr(matrix, row, col-1, nextStr, dir)
			case "-+":
				return checkNextStr(matrix, row-1, col+1, nextStr, dir)
			case "+-":
				return checkNextStr(matrix, row+1, col-1, nextStr, dir)
			}
		}
	}
	return 0
}

func solveProblem2(problemMatrix [][]string) int {
	count := 0
	for i := 0; i < len(problemMatrix); i++ {
		for j := 0; j < len(problemMatrix[i]); j++ {
			if problemMatrix[i][j] == "M" {
				count += checkNextStr2(problemMatrix, i+1, j+1, "A", "++") +
					checkNextStr2(problemMatrix, i-1, j-1, "A", "--") +
					checkNextStr2(problemMatrix, i+1, j-1, "A", "+-") +
					checkNextStr2(problemMatrix, i-1, j+1, "A", "-+")
				problemMatrix[i][j] = "."
			}
		}
	}
	return count
}

func checkNextStr2(matrix [][]string, row int, col int, nextStr string, dir string) int {
	if row >= len(matrix) || row < 0 || col >= len(matrix[row]) || col < 0 {
		return 0
	}

	if matrix[row][col] == nextStr && nextStr != "S" {
		nextStr = "S"
		switch dir {
		case "++":
			return checkNextStr2(matrix, row+1, col+1, nextStr, "++")
		case "--":
			return checkNextStr2(matrix, row-1, col-1, nextStr, "--")
		case "-+":
			return checkNextStr2(matrix, row-1, col+1, nextStr, "-+")
		case "+-":
			return checkNextStr2(matrix, row+1, col-1, nextStr, "+-")
		}
	} else if matrix[row][col] == nextStr {
		r1 := row
		c1 := col
		r2 := row
		c2 := col
		switch dir {
		case "++":
			r1 -= 2
			c2 -= 2
		case "--":
			r1 += 2
			c2 += 2
		case "-+":
			r1 += 2
			c2 -= 2
		case "+-":
			r1 -= 2
			c2 += 2
		}
		if r1 >= len(matrix) || r1 < 0 || c2 >= len(matrix[r1]) || c2 < 0 {
			return 0
		}
		if matrix[r1][c1] == "S" && matrix[r2][c2] == "M" {
			return 1
		} else if matrix[r1][c1] == "M" && matrix[r2][c2] == "S" {
			return 1
		}
	}
	return 0
}
