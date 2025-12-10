package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Col int
	Row int
}

var Directions = []Point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var points []Point
	for scanner.Scan() {
		x, y := 0, 0
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		points = append(points, Point{x, y})
	}

	// fmt.Println(part1(points))
	fmt.Println(part2(points))
}

func part1(points []Point) int {
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x1, x2 := points[i].Col, points[j].Col
			y1, y2 := points[i].Row, points[j].Row

			x := ((x1 - x2) * 2 / 2) + 1
			y := ((y1 - y2) * 2 / 2) + 1

			a := x * y

			if a > maxArea {
				maxArea = a
			}
		}
	}
	return maxArea
}

func part2(points []Point) int {
	grid := BuildGrid(points)
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if CheckGrid(grid, points[i], points[j]) {
				x1, x2 := points[i].Col, points[j].Col
				y1, y2 := points[i].Row, points[j].Row

				x := ((x1 - x2) * 2 / 2) + 1
				y := ((y1 - y2) * 2 / 2) + 1

				a := x * y

				if a > maxArea {
					maxArea = a
				}
			}

		}
	}
	return maxArea
}

func BuildGrid(points []Point) [][]rune {
	knownPoint := Point{8, 2}
	rows, cols := 0, 0
	rowPadding, colPadding := 2, 3
	for _, p := range points {
		if p.Row > rows {
			rows = p.Row
		}

		if p.Col > cols {
			cols = p.Col
		}

	}

	grid := make([][]rune, rows+rowPadding)
	for i := range grid {
		grid[i] = make([]rune, cols+colPadding)
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for i, p := range points {
		grid[p.Row][p.Col] = '#'

		nextI := i + 1
		if nextI == len(points) {
			nextI = 0
		}

		for _, pp := range PointsBetween(p, points[nextI]) {
			grid[pp.Row][pp.Col] = 'X'
		}
	}

	Fill(knownPoint, grid)

	return grid
}

func PointsBetween(p1, p2 Point) []Point {
	var points []Point

	if p1.Row == p2.Row {
		d := 1 // p1 col < p2 col
		if p1.Col > p2.Col {
			d = -1
		}

		for col := p1.Col + d; col != p2.Col; col += d {
			points = append(points, Point{col, p1.Row})
		}

		return points
	} else if p1.Col == p2.Col {
		d := 1 // p1 row < p2 row
		if p1.Row > p2.Row {
			d = -1
		}

		for row := p1.Row + d; row != p2.Row; row += d {
			points = append(points, Point{p1.Col, row})
		}

		return points
	}

	return points
}

func Fill(p Point, grid [][]rune) {
	if grid[p.Row][p.Col] != '.' {
		return
	}

	grid[p.Row][p.Col] = 'X'

	for _, d := range Directions {
		Fill(Point{p.Col + d.Col, p.Row + d.Row}, grid)
	}
}

func CheckGrid(grid [][]rune, p1, p2 Point) bool {
	x1, y1 := p1.Col, p1.Row
	x2, y2 := p2.Col, p2.Row
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			if grid[y][x] == '.' {
				return false
			}
		}
	}
	return true
}
