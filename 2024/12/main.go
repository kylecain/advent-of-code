package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	i int
	j int
}

var up = Point{-1, 0}
var down = Point{1, 0}
var left = Point{0, -1}
var right = Point{0, 1}

var directions = []Point{
	up, down, left, right,
}

var cornerCombinations = [][]Point{
	{up, left},
	{up, right},
	{down, left},
	{down, right},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	plantGrid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		plants := []string{}

		for _, r := range line {
			plants = append(plants, string(r))
		}

		plantGrid = append(plantGrid, plants)
	}

	solve(plantGrid)
}

func solve(grid [][]string) {
	visited := make(map[Point]bool)
	partOne := 0
	partTwo := 0

	for i := range grid {
		for j := range grid {
			p := Point{i, j}
			perimeter, area, corners := 0, 0, 0

			if visited[p] {
				continue
			}
			findFences(p, grid, visited, &perimeter, &area, &corners)
			partOne += perimeter * area
			partTwo += corners * area
		}
	}

	fmt.Println("part one:", partOne)
	fmt.Println("part two:", partTwo)
}

func findFences(p Point, grid [][]string, visited map[Point]bool, perimeter *int, area *int, corners *int) {
	visited[p] = true
	*area++

	for _, c := range cornerCombinations {
		d1, d2 := Point{p.i + c[0].i, p.j + c[0].j}, Point{p.i + c[1].i, p.j + c[1].j}
		if isValidCorner(p, d1, d2, grid) {
			*corners++
		}
	}

	for _, d := range directions {
		p1 := Point{p.i + d.i, p.j + d.j}

		if isValid(p, p1, grid, visited) {
			findFences(p1, grid, visited, perimeter, area, corners)
		} else if !isInBounds(p1, grid) || grid[p.i][p.j] != grid[p1.i][p1.j] {
			*perimeter++
		}
	}
}

func isInBounds(p1 Point, grid [][]string) bool {
	return p1.i >= 0 && p1.j >= 0 && p1.i < len(grid) && p1.j < len(grid[0])
}

func isValid(p Point, p1 Point, grid [][]string, visited map[Point]bool) bool {
	return isInBounds(p1, grid) && !visited[p1] && grid[p.i][p.j] == grid[p1.i][p1.j]
}

func isValidCorner(p Point, d1 Point, d2 Point, grid [][]string) bool {
	isD1InBounds, isD2InBounds := isInBounds(d1, grid), isInBounds(d2, grid)
	if !isD1InBounds && !isD2InBounds {
		return true
	} else if isD1InBounds && isD2InBounds {
		pV, d1V, d2V := grid[p.i][p.j], grid[d1.i][d1.j], grid[d2.i][d2.j]
		if pV != d1V && pV != d2V {
			return true
		} else if pV == d1V && pV == d2V {
			cI, cJ := p.i^d1.i^d2.i, p.j^d1.j^d2.j
			return grid[p.i][p.j] != grid[cI][cJ]
		} else {
			return false
		}
	} else if isD1InBounds && !isD2InBounds {
		return grid[d1.i][d1.j] != grid[p.i][p.j]
	} else if isD2InBounds && !isD1InBounds {
		return grid[d2.i][d2.j] != grid[p.i][p.j]
	}

	return false
}
