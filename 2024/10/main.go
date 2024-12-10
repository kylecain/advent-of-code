package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	trailMap := [][]int{}
	startingPoints := [][]int{}
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		slopes := []int{}

		for i, r := range line {
			slope := int(r - '0')

			if slope == 0 {
				startingPoints = append(startingPoints, []int{i, lineNumber})
			}

			slopes = append(slopes, slope)
		}

		trailMap = append(trailMap, slopes)
		lineNumber++
	}

	for _, v := range trailMap {
		fmt.Println(v)
	}

	partOne(trailMap, startingPoints)
}

func partOne(trailMap [][]int, startingPoints [][]int) {
	total := 0

	rows := len(trailMap)
	cols := len(trailMap[0])
	visitedPoints := make([][]bool, rows)
	visitedNines := make(map[Point]bool)

	part1sum := 0

	for i := range visitedPoints {
		visitedPoints[i] = make([]bool, cols)
	}

	for _, point := range startingPoints {
		fmt.Println("starting point", point)
		findTrail(trailMap, visitedPoints, point[0], point[1], &total, visitedNines)
		part1sum += len(visitedNines)
		visitedNines = make(map[Point]bool)
	}

	fmt.Println(total)
	fmt.Println(part1sum)
}

type Point struct {
	y int
	x int
}

func findTrail(trailMap [][]int, visitedPoints [][]bool, x int, y int, total *int, visitedNines map[Point]bool) {
	// visitedPoints[y][x] = true

	if trailMap[y][x] == 9 {
		fmt.Println("reached 9", x, y)
		*total++
		visitedNines[Point{y, x}] = true
	}

	for _, direction := range directions {
		x1, y1 := x+direction[0], y+direction[1]

		if isValid(trailMap, visitedPoints, x1, y1, trailMap[y][x]) {
			fmt.Println(x, y, trailMap[y][x])
			findTrail(trailMap, visitedPoints, x1, y1, total, visitedNines)
		}
	}
}

func isValid(trailMap [][]int, visitedPoints [][]bool, x int, y int, previousValue int) bool {
	isInBounds := x >= 0 && y >= 0 && x < len(trailMap[0]) && y < len(trailMap)
	return isInBounds && !visitedPoints[y][x] && (trailMap[y][x]-previousValue == 1)
}
