package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var list1 []int
	var list2 []int
	var err error
	var line []byte
	for {
		line, _, err = reader.ReadLine()
		if err != nil || string(line) == "" {
			break
		}
		values := strings.Split(string(line), "   ")
		v1, _ := strconv.Atoi(values[0])
		v2, _ := strconv.Atoi(values[1])
		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	problem1(list1, list2)
	problem2(list1, list2)

}

func problem1(list1 []int, list2 []int) {
	totalSize := 0.
	for i := 0; i < len(list1); i++ {
		totalSize += math.Abs(float64(list2[i] - list1[i]))
	}
	fmt.Println("Total size: ", totalSize)
}
func problem2(list1 []int, list2 []int) {
	ptr1 := 0
	ptr2 := 0
	timesForEach := make(map[int]int, 0)
	for ptr1 < len(list1) && ptr2 < len(list2) {
		if _, ok := timesForEach[list1[ptr1]]; !ok {
			timesForEach[list1[ptr1]] = 0
		}

		if list1[ptr1] > list2[ptr2] {
			ptr2++
		} else if list1[ptr1] < list2[ptr2] {
			ptr1++
		} else {
			timesForEach[list1[ptr1]] += 1
			ptr2++
		}
	}
	totalDistance := 0
	for _, v := range list1 {
		totalDistance += (v * timesForEach[v])
	}
	fmt.Println("Total distance: ", totalDistance)

}
