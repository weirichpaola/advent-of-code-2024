package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var err error
	var line []byte
	reports := [][]string{}
	for {
		line, _, err = reader.ReadLine()
		if err != nil || string(line) == "" {
			break
		}
		values := strings.Split(string(line), " ")
		reports = append(reports, values)
	}

	numberSafeReports := problem1(reports)
	fmt.Println("Safe reports ", numberSafeReports)
	numberSafeReports = problem2(reports)
	fmt.Println("Safe reports ", numberSafeReports)

}

func problem1(reports [][]string) int {
	safeReports := 0
	for _, report := range reports {
		if isValidReport(report) {
			safeReports++
		}
	}
	return safeReports
}

func problem2(reports [][]string) int {
	safeReports := 0
	for _, report := range reports {
		if isValidReport(report) {
			safeReports++
		} else {
			hasSafeRemovingOne := false
			for i, _ := range report {
				test := make([]string, len(report))
				copy(test, report)
				if isValidReport(append(test[:i], test[i+1:]...)) {
					hasSafeRemovingOne = true
					break
				}
			}
			if hasSafeRemovingOne {
				safeReports++
			}
		}
	}
	return safeReports
}

func isValidReport(report []string) bool {
	f, _ := strconv.Atoi(report[0])
	s, _ := strconv.Atoi(report[1])
	isDecreasing := f > s
	isValid := f != s && math.Abs(float64(f-s)) < 4
	if !isValid {
		return false
	}
	prevReport := s
	for i := 2; i < len(report); i++ {
		currReport, _ := strconv.Atoi(report[i])
		if isDecreasing {
			if prevReport <= currReport || math.Abs(float64(prevReport-currReport)) > 3 {
				return false
			}
		} else {
			if prevReport >= currReport || math.Abs(float64(currReport-prevReport)) > 3 {
				return false
			}
		}

		prevReport = currReport
	}
	return true
}
