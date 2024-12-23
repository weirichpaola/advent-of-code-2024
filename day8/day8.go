package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./day8/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	antennas := [][]string{}
	antinodes := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		valuesStr := strings.Split(line, "")
		values := []string{}
		for _, v := range valuesStr {
			values = append(values, v)
		}
		antennas = append(antennas, values)
		antinodes = append(antinodes, values)
	}

	res := problem1(antennas, copyMap(antinodes))
	print(res)
	fmt.Println("\n")
	res2 := problem2(antennas, copyMap(antinodes))
	print(res2)
}

func problem1(antennas [][]string, antinodes [][]string) int {
	pos := map[string][][]int{}
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[0]); j++ {
			if antennas[i][j] != "." {
				v := antennas[i][j]
				if _, ok := pos[v]; ok {
					pos[v] = append(pos[v], []int{i, j})
				} else {
					pos[v] = [][]int{{i, j}}
				}
			}
		}
	}

	for _, letter := range pos {
		for i := 0; i < len(letter)-1; i++ {
			x1 := letter[i][0]
			y1 := letter[i][1]
			for j := i + 1; j < len(letter); j++ {
				x2 := letter[j][0]
				y2 := letter[j][1]
				posx := 0
				posx2 := 0
				posy := 0
				posy2 := 0
				if x1 < x2 {
					posx = x1 - (x2 - x1)
					posx2 = x2 + (x2 - x1)
				} else {
					posx = x1 + (x1 - x2)
					posx2 = x2 - (x1 - x2)
				}
				if y1 < y2 {
					posy = y1 - (y2 - y1)
					posy2 = y2 + (y2 - y1)
				} else {
					posy = y1 + (y1 - y2)
					posy2 = y2 - (y1 - y2)
				}
				if posx >= 0 && posx < len(antinodes) && posy >= 0 && posy < len(antinodes[0]) {
					antinodes[posx][posy] = "#"
				}
				if posx2 >= 0 && posx2 < len(antinodes) && posy2 >= 0 && posy2 < len(antinodes[0]) {
					antinodes[posx2][posy2] = "#"
				}
			}
		}
	}
	numAntinodes := 0
	for i := 0; i < len(antinodes); i++ {
		for j := 0; j < len(antinodes[0]); j++ {
			if antinodes[i][j] == "#" {
				numAntinodes++
			}
			fmt.Print(antinodes[i][j], " ")
		}
		fmt.Print("\n")
	}
	return numAntinodes
}

func problem2(antennas [][]string, antinodes [][]string) int {
	pos := map[string][][]int{}
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[0]); j++ {
			if antennas[i][j] != "." {
				v := antennas[i][j]
				if _, ok := pos[v]; ok {
					pos[v] = append(pos[v], []int{i, j})
				} else {
					pos[v] = [][]int{{i, j}}
				}
			}
		}
	}

	for _, letter := range pos {
		for i := 0; i < len(letter)-1; i++ {
			for j := i + 1; j < len(letter); j++ {
				x1 := letter[i][0]
				y1 := letter[i][1]
				x2 := letter[j][0]
				y2 := letter[j][1]

				posx := 0
				posx2 := 0
				posy := 0
				posy2 := 0
				isNextHashPossible := true
				diffx := int(math.Abs(float64(x2 - x1)))
				diffy := int(math.Abs(float64(y2 - y1)))

				for isNextHashPossible {
					if x1 < x2 {
						posx = x1 - diffx
						posx2 = x2 + diffx
					} else {
						posx = x1 + diffx
						posx2 = x2 - diffx
					}
					if y1 < y2 {
						posy = y1 - diffy
						posy2 = y2 + diffy
					} else {
						posy = y1 + diffy
						posy2 = y2 - diffy
					}
					isOnePossible := true
					isTwoPossible := true
					if posx >= 0 && posx < len(antinodes) && posy >= 0 && posy < len(antinodes[0]) {
						antinodes[posx][posy] = "#"
						x1 = posx
						y1 = posy
					} else {
						isOnePossible = false
					}
					if posx2 >= 0 && posx2 < len(antinodes) && posy2 >= 0 && posy2 < len(antinodes[0]) {
						antinodes[posx2][posy2] = "#"
						x2 = posx2
						y2 = posy2
					} else {
						isTwoPossible = false
					}
					if !isOnePossible && !isTwoPossible {
						isNextHashPossible = false
					}
				}
			}
		}
	}
	numAntinodes := 0
	for i := 0; i < len(antinodes); i++ {
		for j := 0; j < len(antinodes[0]); j++ {
			if antinodes[i][j] != "." {
				numAntinodes++
			}
			fmt.Print(antinodes[i][j], " ")
		}
		fmt.Print("\n")
	}
	return numAntinodes
}

func copyMap(currentMap [][]string) [][]string {
	duplicate := make([][]string, len(currentMap))
	for i := range currentMap {
		duplicate[i] = make([]string, len(currentMap[i]))
		copy(duplicate[i], currentMap[i])
	}
	return duplicate
}
