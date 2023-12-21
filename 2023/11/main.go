package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var universe []string

	for scanner.Scan() {
		universe = append(universe, scanner.Text())
	}

	var expansionIndicies []int
	for i, s := range universe {
		hasGalaxy := false
		for _, r := range s {
			if string(r) == "#" {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			expansionIndicies = append(expansionIndicies, i)
		}
	}

	expansionString := strings.Repeat(".", len(universe[0]))

	for i, j := range expansionIndicies {
		universe = append(universe[:j+i], append([]string{expansionString}, universe[j+i:]...)...)
	}

	expansionIndicies = []int{}

	for j, _ := range universe[0] {
		hasGalaxy := false
		for i, _ := range universe {
			if string(universe[i][j]) == "#" {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			expansionIndicies = append(expansionIndicies, j)
		}
	}

	for i, j := range expansionIndicies {
		for k := range universe {
			universe[k] = universe[k][:j+i] + "." + universe[k][j+i:]
		}
	}

	points := []Point{}
	for i, str := range universe {
		for j, r := range str {
			if string(r) == "#" {
				points = append(points, Point{X: j, Y: i})
			}
		}
	}

	sum := 0
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			sum += findShortestPath(universe, point, points[j])
		}
	}
	fmt.Println(sum)
}

func findShortestPath(universe []string, start Point, end Point) int {
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	queue := []struct {
		point  Point
		length int
	}{{start, 0}}
	visited := make(map[Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == end {
			return current.length
		}

		for _, direction := range directions {
			x := current.point.X + direction[0]
			y := current.point.Y + direction[1]

			if x >= 0 && x < len(universe[0]) && y >= 0 && y < len(universe) && !visited[Point{x, y}] {
				visited[Point{x, y}] = true

				queue = append(queue, struct {
					point  Point
					length int
				}{Point{x, y}, current.length + 1})
			}
		}
	}

	return -1
}
