package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Pair = Direction

type Direction struct {
	X int
	Y int
}

var left = Direction{-1, 0}
var right = Direction{1, 0}
var up = Direction{0, -1}
var down = Direction{0, 1}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y, lineNumber := 0, 0, 0
	labMap := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		exp := regexp.MustCompile(`\^`)
		startingPosition := exp.FindStringIndex(line)

		if len(startingPosition) != 0 {
			x = startingPosition[0]
			y = lineNumber
		}

		labMap = append(labMap, []rune(line))
		lineNumber++
	}

	nextDirections := map[Direction]Direction{
		up:    right,
		right: down,
		down:  left,
		left:  up,
	}

	partOne(labMap, nextDirections, x, y)
	partTwo(labMap, nextDirections, x, y)
}

func partOne(labMap [][]rune, nextDirections map[Direction]Direction, x, y int) {
	visitedTiles := make(map[Pair]int)

	isSearching := true
	direction := up

	for isSearching {
		visitedTiles[Pair{x, y}]++

		xCanidate, yCanidate := x+direction.X, y+direction.Y
		if !isInBounds(labMap, xCanidate, yCanidate) {
			isSearching = false
			break
		}

		if labMap[yCanidate][xCanidate] == '#' {
			direction = nextDirections[direction]
			xCanidate, yCanidate = x+direction.X, y+direction.Y
		}

		x, y = xCanidate, yCanidate
	}

	fmt.Println(len(visitedTiles))
}

func partTwo(labMap [][]rune, nextDirections map[Direction]Direction, x, y int) {
	barriersSum := 0

	for i, v := range labMap {
		for j := range v {
			if j == x && i == y {
				continue
			}

            lm := make([][]rune, len(labMap))
            for i := range labMap {
                lm[i] = make([]rune, len(labMap[i]))
                copy(lm[i], labMap[i])
            }

			lm[i][j] = '#'

			barriersSum += isSolveable(lm, nextDirections, x, y)
		}
	}

	fmt.Println(barriersSum)
}

func isSolveable(labMap [][]rune, nextDirections map[Direction]Direction, x, y int) int {
	direction := up
	hitBarriers := make(map[Pair]int)

	for true {
		xCanidate, yCanidate := x + direction.X, y + direction.Y
		if !isInBounds(labMap, xCanidate, yCanidate) {
			break
		}

		if labMap[yCanidate][xCanidate] == '#' {
            hitBarriers[Pair{xCanidate, yCanidate}]++

            if hitBarriers[Pair{xCanidate, yCanidate}] > 10 {
                return 1
            }

			direction = nextDirections[direction]
			xCanidate, yCanidate = x + direction.X, y + direction.Y

            if labMap[yCanidate][xCanidate] == '#' {
                direction = nextDirections[direction]
                xCanidate, yCanidate = x + direction.X, y + direction.Y
            }
		}

		x, y = xCanidate, yCanidate
	}

	return 0
}

func isInBounds(grid [][]rune, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid)
}
