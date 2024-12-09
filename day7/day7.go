package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./day7/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	expectedResultsList := []int{}
	valuesList := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		numbers := strings.Split(line, ":")
		expectedResult := numbers[0]
		r, _ := strconv.Atoi(expectedResult)
		expectedResultsList = append(expectedResultsList, r)
		valuesStr := strings.Split(strings.TrimPrefix(numbers[1], " "), " ")
		values := []int{}
		for _, v := range valuesStr {
			r, _ = strconv.Atoi(v)
			values = append(values, r)
		}
		valuesList = append(valuesList, values)
	}

	res := problem1(expectedResultsList, valuesList)
	fmt.Println(res)

	res2 := problem2(expectedResultsList, valuesList)
	fmt.Println(res2)
}

func problem1(expectedResultsList []int, valuesList [][]int) int {
	finalRes := 0
	for i, _ := range expectedResultsList {
		if isValid(expectedResultsList[i], valuesList[i][0], valuesList[i], "*", 0) || isValid(expectedResultsList[i], valuesList[i][0], valuesList[i], "+", 0) {
			finalRes += expectedResultsList[i]
		}
	}
	return finalRes
}

func isValid(expectedResult int, currVal int, values []int, op string, pos int) bool {
	if pos == len(values)-1 {
		return false
	}

	switch op {
	case "*":
		currVal *= values[pos+1]
		if currVal > expectedResult {
			return false
		} else if currVal == expectedResult && pos == len(values)-2 {
			return true
		}
	case "+":
		currVal += values[pos+1]
		if currVal > expectedResult {
			return false
		} else if currVal == expectedResult && pos == len(values)-2 {
			return true
		}
	}
	return isValid(expectedResult, currVal, values, "*", pos+1) || isValid(expectedResult, currVal, values, "+", pos+1)

}

func problem2(expectedResultsList []int, valuesList [][]int) int {
	finalRes := 0
	for i, _ := range expectedResultsList {
		if isValid2(expectedResultsList[i], valuesList[i][0], valuesList[i], "*", 0) || isValid2(expectedResultsList[i], valuesList[i][0], valuesList[i], "+", 0) || isValid2(expectedResultsList[i], valuesList[i][0], valuesList[i], "||", 0) {
			finalRes += expectedResultsList[i]
		}
	}
	return finalRes
}

func isValid2(expectedResult int, currVal int, values []int, op string, pos int) bool {
	if pos == len(values)-1 {
		return false
	}

	switch op {
	case "*":
		currVal *= values[pos+1]
		if currVal > expectedResult {
			return false
		} else if currVal == expectedResult && pos == len(values)-2 {
			return true
		}
	case "+":
		currVal += values[pos+1]
		if currVal > expectedResult {
			return false
		} else if currVal == expectedResult && pos == len(values)-2 {
			return true
		}
	case "||":
		currVal, _ = strconv.Atoi(fmt.Sprintf("%d%d", currVal, values[pos+1]))
		if currVal > expectedResult {
			return false
		} else if currVal == expectedResult && pos == len(values)-2 {
			return true
		}
	}
	return isValid2(expectedResult, currVal, values, "*", pos+1) || isValid2(expectedResult, currVal, values, "+", pos+1) || isValid2(expectedResult, currVal, values, "||", pos+1)

}
