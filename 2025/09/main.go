package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type Point struct {
	Col int
	Row int
}

type Combination struct {
	p1, p2 Point
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

	fmt.Println(part1(points))
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
	n := len(points)
	work := make(chan Combination, 1000)
	go func() {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				work <- Combination{points[i], points[j]}
			}
		}
		close(work)
	}()

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	results := make(chan int, numWorkers)

	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			localMax := 0
			for combination := range work {
				if CheckGrid(grid, combination.p1, combination.p2) {
					x := abs(combination.p1.Col-combination.p2.Col) + 1
					y := abs(combination.p1.Row-combination.p2.Row) + 1
					a := x * y
					if a > localMax {
						localMax = a
					}
				}
			}
			results <- localMax
		}()
	}

	wg.Wait()
	close(results)

	maxArea := 0
	for a := range results {
		if a > maxArea {
			maxArea = a
		}
	}
	return maxArea
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func PrintBitGrid(grid [][]uint64, cols int) {
	for _, row := range grid {
		for c := 0; c < cols; c++ {
			word := c / 64
			bit := uint(c % 64)
			if (row[word] & (1 << bit)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}

func wordIndex(col int) (wordIdx int, bit uint) {
	wordIdx = col / 64
	bit = uint(col % 64)
	return
}

func Set(grid [][]uint64, row, col int, value bool) {
	w, b := wordIndex(col)
	if value {
		grid[row][w] |= 1 << b
	} else {
		grid[row][w] &^= 1 << b
	}
}

func Get(grid [][]uint64, row, col int) bool {
	w, b := wordIndex(col)
	return (grid[row][w] & (1 << b)) != 0
}

func setRange(grid [][]uint64, row, start, end int) {
	if start > end {
		return
	}

	sw, sb := wordIndex(start)
	ew, eb := wordIndex(end)

	if sw == ew {
		length := uint(eb - sb + 1)
		mask := ((uint64(1) << length) - 1) << sb
		grid[row][sw] |= mask
		return
	}

	startMask := ^uint64(0) << sb
	grid[row][sw] |= startMask

	for w := sw + 1; w <= ew-1; w++ {
		grid[row][w] = ^uint64(0)
	}

	endMask := (uint64(1) << uint(eb+1)) - 1
	grid[row][ew] |= endMask
}

func BuildGrid(points []Point) [][]uint64 {
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

	words := (cols + 63) / 64
	grid := make([][]uint64, rows)
	for i := range grid {
		grid[i] = make([]uint64, words)
	}

	for i, p := range points {
		Set(grid, p.Row, p.Col, true)

		nextI := i + 1
		if nextI == len(points) {
			nextI = 0
		}

		for _, pp := range PointsBetween(p, points[nextI]) {
			Set(grid, pp.Row, pp.Col, true)
		}
	}

	Fill(grid, cols)

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

func Fill(grid [][]uint64, cols int) {
	nRows := len(grid)
	if nRows == 0 || cols <= 0 {
		return
	}

	numWorkers := runtime.NumCPU()
	chunkSize := (nRows + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		startRow := w * chunkSize
		endRow := min(startRow+chunkSize, nRows)

		go func(start, end, workerID int) {
			defer wg.Done()
			chunkRows := end - start
			logInterval := chunkRows / 100
			if logInterval == 0 {
				logInterval = 1
			}

			for row := start; row < end; row++ {
				var edges []int
				prev := false
				for col := 0; col < cols; col++ {
					cur := Get(grid, row, col)
					if cur && !prev {
						edges = append(edges, col)
					}
					prev = cur
				}

				if len(edges) >= 2 {
					for i := 0; i+1 < len(edges); i += 2 {
						col1, col2 := edges[i], edges[i+1]
						startCol := col1 + 1
						endCol := col2 - 1
						if startCol <= endCol {
							setRange(grid, row, startCol, endCol)
						}
					}
				}
			}
		}(startRow, endRow, w)
	}

	wg.Wait()
}

func CheckGrid(grid [][]uint64, p1, p2 Point) bool {
	row1, col1 := p1.Col, p1.Row
	row2, col2 := p2.Col, p2.Row

	if row1 > row2 {
		row1, row2 = row2, row1
	}
	if col1 > col2 {
		col1, col2 = col2, col1
	}

	for row := col1; row <= col2; row++ {
		startWord, startBit := wordIndex(row1)
		endWord, endBit := wordIndex(row2)

		for w := startWord; w <= endWord; w++ {
			word := grid[row][w]
			mask := ^uint64(0)

			if w == startWord {
				mask &= ^uint64(0) << startBit
			}

			if w == endWord {
				mask &= (uint64(1) << (endBit + 1)) - 1
			}

			if word&mask != mask {
				return false
			}
		}
	}

	return true
}
