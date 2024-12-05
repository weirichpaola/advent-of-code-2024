package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/I860770/Documents/adventOfCode2024/day3/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	problemStr := ""
	for scanner.Scan() {
		line := scanner.Text()
		problemStr += line
	}
	resultProblem1 := problem1Check(problemStr)
	resultProblem2 := problem2Check(problemStr)
	fmt.Println("Result problem 1: ", resultProblem1)
	fmt.Println("Result problem 2: ", resultProblem2)
}

func problem2Check(line string) int {
	newValues := strings.Split(line, "do()")
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	resultProblem2 := 0
	for _, v := range newValues {
		if strings.Contains(v, "don't()") {
			v = strings.SplitAfter(v, "don't()")[0]
			v = strings.TrimPrefix(v, "don't()")
		}
		if v != "" {
			multValues := re.FindAllString(v, -1)
			for _, v := range multValues {
				newData := strings.TrimSuffix(strings.TrimPrefix(v, "mul("), ")")
				values := strings.Split(newData, ",")
				v1, _ := strconv.Atoi(values[0])
				v2, _ := strconv.Atoi(values[1])

				resultProblem2 += v1 * v2
			}
		}
	}
	return resultProblem2
}

func problem1Check(line string) int {
	resultProblem1 := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	multValues := re.FindAllString(line, -1)
	for _, v := range multValues {
		newData := strings.TrimSuffix(strings.TrimPrefix(v, "mul("), ")")
		values := strings.Split(newData, ",")
		v1, _ := strconv.Atoi(values[0])
		v2, _ := strconv.Atoi(values[1])

		resultProblem1 += v1 * v2
	}
	return resultProblem1
}
