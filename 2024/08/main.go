package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Coordinate struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	antennaMap := make(map[rune][]Coordinate)
	antennaGrid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		exp := regexp.MustCompile(`[^\.]`)
		antennas := exp.FindAllStringIndex(line, -1)

		for _, antenna := range antennas {
			antennaRune := rune(line[antenna[0]])
			antennaCoordinate := Coordinate{antenna[0], lineNumber}

			antennaMap[antennaRune] = append(antennaMap[antennaRune], antennaCoordinate)
		}

		antennaGrid = append(antennaGrid, []rune(line))
		lineNumber++
	}

	// partOne(antennaMap, antennaGrid)
	partTwo(antennaMap, antennaGrid)
}

func partOne(antennaMap map[rune][]Coordinate, antennaGrid [][]rune) {
	antinodeMap := make(map[Coordinate]bool)

	for _, coordinates := range antennaMap {
		for i := 0; i < len(coordinates); i++ {
			for j := i + 1; j < len(coordinates); j++ {
				combination := []Coordinate{coordinates[j], coordinates[i]}

				p1, p2 := combination[0], combination[1]

				x := 2*p1.X - p2.X
				y := 2*p1.Y - p2.Y

				antiNode := Coordinate{x, y}

				if isInBounds(antennaGrid, x, y) {
					antinodeMap[antiNode] = true
				}

				x = 2*p2.X - p1.X
				y = 2*p2.Y - p1.Y

				antiNode = Coordinate{x, y}

				if isInBounds(antennaGrid, x, y) {
					antinodeMap[antiNode] = true
				}
			}
		}
	}

	fmt.Println(len(antinodeMap))
}

func partTwo(antennaMap map[rune][]Coordinate, antennaGrid [][]rune) {
	antinodeMap := make(map[Coordinate]bool)

	for _, coordinates := range antennaMap {
		for i := 0; i < len(coordinates); i++ {
			for j := i + 1; j < len(coordinates); j++ {
				combination := []Coordinate{coordinates[j], coordinates[i]}

				p1, p2 := combination[0], combination[1]

				antinodeMap[p1] = true
				antinodeMap[p2] = true

				addAntinodes(antinodeMap, antennaGrid, p1, p2)
				addAntinodes(antinodeMap, antennaGrid, p2, p1)
			}
		}
	}

	fmt.Println(len(antinodeMap))
}

func addAntinodes(
	antinodeMap map[Coordinate]bool,
	antennaGrid [][]rune,
	p1 Coordinate,
	p2 Coordinate,
) {
	x := 2*p1.X - p2.X
	y := 2*p1.Y - p2.Y

	antinode := Coordinate{x, y}

	if isInBounds(antennaGrid, x, y) {
		antinodeMap[antinode] = true
		addAntinodes(antinodeMap, antennaGrid, antinode, p1)
	} else {
		return
	}
}

func isInBounds(grid [][]rune, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid)
}
