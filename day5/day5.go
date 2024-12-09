package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day5/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	afterNumber := make(map[int][]int, 0)
	currentUpdate := [][]int{}
	currentUpdateHappening := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentUpdateHappening = true
			continue
		}
		if currentUpdateHappening {
			updateValues := strings.Split(line, ",")
			values := []int{}
			for _, up := range updateValues {
				v, _ := strconv.Atoi(up)
				values = append(values, v)
			}
			currentUpdate = append(currentUpdate, values)
		} else {
			values := strings.Split(line, "|")
			v1, _ := strconv.Atoi(values[0])
			v2, _ := strconv.Atoi(values[1])
			if _, ok := afterNumber[v1]; !ok {
				afterNumber[v1] = []int{v2}
			} else {
				afterNumber[v1] = append(afterNumber[v1], v2)
			}
		}

	}

	problem1 := solveProblem1(afterNumber, currentUpdate)
	fmt.Println(problem1)

	problem2 := solveProblem2(afterNumber, currentUpdate)
	fmt.Println(problem2)
}

func solveProblem1(afterNumber map[int][]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		foundInvalid := false
		foundNumbers := make(map[int]struct{}, 0)
		for _, v := range update {
			afterNumber := afterNumber[v]
			for _, af := range afterNumber {
				if _, ok := foundNumbers[af]; ok {
					foundInvalid = true
					break
				}
			}
			foundNumbers[v] = struct{}{}
			if foundInvalid {
				break
			}
		}
		if !foundInvalid {
			sum += update[int(len(update)/2)]
		}
	}
	return sum
}

func solveProblem2(afterNumber map[int][]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		foundInvalid := false
		foundNumbers := make(map[int]int, 0)
		for pos, v := range update {
			afterNumber := afterNumber[v]
			for _, af := range afterNumber {
				if _, ok := foundNumbers[af]; ok {
					foundInvalid = true
				}
			}
			foundNumbers[v] = pos
		}
		if foundInvalid {
			sort.Slice(update, func(i, j int) bool {
				return isValid(update[i], update[j], afterNumber)
			})
			sum += update[int(len(update)/2)]
		}
	}
	return sum
}

func isValid(num1 int, num2 int, afterNumber map[int][]int) bool {
	isValid := true
	for _, n := range afterNumber[num2] {
		if n == num1 {
			isValid = false
		}
	}
	return isValid
}
