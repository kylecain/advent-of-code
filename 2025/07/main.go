package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func (p *Point) canSpawnLeft() bool {
	return p.Col-1 >= 0
}

func (p *Point) canSpawnRight(m []string) bool {
	return p.Col+1 < len(m[0])
}

func (p *Point) canSpawnDown(m []string) bool {
	return p.Row+1 < len(m)
}

func (p *Point) NewLeft() Point {
	return Point{p.Row, p.Col - 1}
}

func (p *Point) NewRight() Point {
	return Point{p.Row, p.Col + 1}
}

func (p *Point) NewDown() Point {
	return Point{p.Row + 1, p.Col}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var manifold []string
	var start Point
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		si := strings.Index(line, "S")
		if si != -1 {
			start = Point{i, si}
		}

		manifold = append(manifold, line)
		i++
	}

	splitters := make(map[Point]bool)
	part1(manifold, start, splitters)
	fmt.Println(len(splitters))

	fmt.Println(part2(manifold, start))
}

func part1(manifold []string, start Point, splitters map[Point]bool) {
	if !start.canSpawnDown(manifold) {
		return
	}

	dp := start.NewDown()

	if string(manifold[dp.Row][dp.Col]) == "^" {
		if !splitters[dp] {
			splitters[dp] = true

			if dp.canSpawnLeft() {
				part1(manifold, dp.NewLeft(), splitters)
			}
			if dp.canSpawnRight(manifold) {
				part1(manifold, dp.NewRight(), splitters)
			}
		}
	} else {
		part1(manifold, dp, splitters)
	}
}

func part2(manifold []string, start Point) int {
	memo := make([]int, len(manifold[0]))
	memo[start.Col] = 1

	for _, line := range manifold {
		for col, char := range line {
			if string(char) == "^" {
				if col-1 >= 0 {
					memo[col-1] += memo[col]
				}
				if col+1 < len(memo) {
					memo[col+1] += memo[col]
				}
				memo[col] = 0
			}
		}
	}

	sum := 0
	for _, v := range memo {
		sum += v
	}

	return sum
}
