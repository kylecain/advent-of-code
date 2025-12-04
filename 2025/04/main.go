package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	J int
	I int
}

var directions = []Point{
	{-1, 0},  // left
	{1, 0},   // right
	{0, -1},  // up
	{0, 1},   // down
	{-1, -1}, // up-left
	{1, -1},  // up-right
	{-1, 1},  // down-left
	{1, 1},   // down-right
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1(lines)
	part2(lines)
}

func part1(grid []string) {
	c := 0
	for i, rolls := range grid {
		for j, roll := range rolls {
			if string(roll) == "@" {
				cc := 0
				for _, d := range directions {
					jj, ii := j+d.J, i+d.I
					if isInbounds(grid, jj, ii) {
						if string(grid[ii][jj]) == "@" {
							cc++
						}
					}
				}
				if cc < 4 {
					c++
				}
			}
		}
	}
	fmt.Println(c)
}

func part2(grid []string) {
	c := 0
	rollsToRemove := []Point{}
	for {
		for i, rolls := range grid {
			for j, roll := range rolls {
				if string(roll) == "@" {
					cc := 0
					for _, d := range directions {
						jj, ii := j+d.J, i+d.I
						if isInbounds(grid, jj, ii) {
							if string(grid[ii][jj]) == "@" {
								cc++
							}
						}
					}
					if cc < 4 {
						c++
						rollsToRemove = append(rollsToRemove, Point{j, i})
					}
				}
			}
		}

		for _, roll := range rollsToRemove {
			str := []rune(grid[roll.I])
			str[roll.J] = 'X'
			grid[roll.I] = string(str)
		}

		if len(rollsToRemove) == 0 {
			break
		}

		rollsToRemove = []Point{}
	}
	fmt.Println(c)
}

func isInbounds(grid []string, j, i int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
