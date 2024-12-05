package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	X int
	Y int
}

var left = Direction{-1, 0}
var right = Direction{1, 0}
var up = Direction{0, 1}
var down = Direction{0, -1}
var upLeft = Direction{-1, 1}
var upRight = Direction{1, 1}
var downLeft = Direction{-1, -1}
var downRight = Direction{1, -1}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	letterGrid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		letterGrid = append(letterGrid, []rune(line))
	}

	partOne(letterGrid)
	partTwo(letterGrid)
}

func partOne(letterGrid [][]rune) {
	directions := []Direction{
		left, right, up, down, upLeft, upRight, downLeft, downRight,
	}

	sum := 0

	for y, runes := range letterGrid {
		for x, rune := range runes {
			if rune == 'X' {
				for _, direction := range directions {
					if digForChristmas(letterGrid, x, y, direction) {
						sum++
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func partTwo(letterGrid [][]rune) {
	sum := 0

	for i, runes := range letterGrid {
		for j, r := range runes {
			runeMap := make(map[rune]int)

			var tl rune
			var tr rune

			if r == 'A' {
				if isInBounds(letterGrid, j+upLeft.X, i+upLeft.Y) {
					subject := letterGrid[i+upLeft.Y][j+upLeft.X]
					runeMap[subject]++

					tl = subject
				}

				if isInBounds(letterGrid, j+upRight.X, i+upRight.Y) {
					subject := letterGrid[i+upRight.Y][j+upRight.X]
					runeMap[subject]++

					tr = subject
				}

				if isInBounds(letterGrid, j+downLeft.X, i+downLeft.Y) {
					subject := letterGrid[i+downLeft.Y][j+downLeft.X]
					runeMap[subject]++

					if tr == subject {
						continue
					}
				}

				if isInBounds(letterGrid, j+downRight.X, i+downRight.Y) {
					subject := letterGrid[i+downRight.Y][j+downRight.X]
					runeMap[subject]++

					if tl == subject {
						continue
					}
				}

				if runeMap['M'] == 2 && runeMap['S'] == 2 {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func isInBounds(letterGrid [][]rune, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(letterGrid[0]) && y < len(letterGrid)
}

func digForChristmas(letterGrid [][]rune, x int, y int, direction Direction) bool {
	target := []rune{'X', 'M', 'A', 'S'}

	for _, v := range target {
		if !isInBounds(letterGrid, x, y) {
			return false
		}

		if letterGrid[y][x] == v {
			x = x + direction.X
			y = y + direction.Y
		} else {
			return false
		}
	}

	return true
}
