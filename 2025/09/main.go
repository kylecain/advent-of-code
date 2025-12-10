package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func BuildGrid(points []Point) [][]bool {
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
	rows += rowPadding
	cols += colPadding

	// knownPoint := Point{cols / 2, rows / 3}
	// knownPoint := Point{9, 4}
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = false
		}
	}

	for i, p := range points {
		grid[p.Row][p.Col] = true

		nextI := i + 1
		if nextI == len(points) {
			nextI = 0
		}

		for _, pp := range PointsBetween(p, points[nextI]) {
			grid[pp.Row][pp.Col] = true
		}
	}

	Fill(grid)

	return grid
}

func PointsBetween(p1, p2 Point) []Point {
	var points []Point

	if p1.Row == p2.Row {
		d := 1
		if p1.Col > p2.Col {
			d = -1
		}

		for col := p1.Col + d; col != p2.Col; col += d {
			points = append(points, Point{col, p1.Row})
		}

		return points
	}

	if p1.Col == p2.Col {
		d := 1
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

func Fill(grid [][]bool) {
	for row := 0; row < len(grid); row++ {
		var cols []int
		prev := false
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] && !prev {
				cols = append(cols, col)
			}
			prev = grid[row][col]
		}

		if len(cols) < 2 {
			continue
		}

		for i := 0; i+1 < len(cols); i += 2 {
			col1, col2 := cols[i], cols[i+1]
			for col := col1 + 1; col < col2; col++ {
				if !grid[row][col] {
					grid[row][col] = true
				}
			}
		}
	}
}

func CheckGrid(grid [][]bool, p1, p2 Point) bool {
	row1, col1 := p1.Col, p1.Row
	row2, col2 := p2.Col, p2.Row
	if row1 > row2 {
		row1, row2 = row2, row1
	}
	if col1 > col2 {
		col1, col2 = col2, col1
	}

	for row := col1; row <= col2; row++ {
		for col := row1; col <= row2; col++ {
			if !grid[row][col] {
				return false
			}
		}
	}
	return true
}

func PrintGrid(grid [][]bool) {
	for _, line := range grid {
		for _, v := range line {
			if v {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
