package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	verticalIndicies := []int{}
	horizontalIndicies := []int{}
	for i, s := range universe {
		hasGalaxy := false
		for _, r := range s {
			if string(r) == "#" {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			verticalIndicies = append(verticalIndicies, i)
		}
	}

	for j, _ := range universe[0] {
		hasGalaxy := false
		for i, _ := range universe {
			if string(universe[i][j]) == "#" {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			horizontalIndicies = append(horizontalIndicies, j)
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

	sum := 0.0
	fmt.Println(horizontalIndicies)
	fmt.Println(verticalIndicies)
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			x1, x2 := point.X, points[j].X
			y1, y2 := point.Y, points[j].Y
			x1Add, x2Add, y1Add, y2Add := 0, 0, 0, 0
			expandBy := 1000000 - 1

			for _, ind := range horizontalIndicies {
				if ind < x1 {
					x1Add += expandBy
				}
				if ind < x2 {
					x2Add += expandBy
				}
			}
			for _, ind := range verticalIndicies {
				if ind < y1 {
					y1Add += expandBy
				}
				if ind < y2 {
					y2Add += expandBy
				}
			}

			x1 += x1Add
			x2 += x2Add
			y1 += y1Add
			y2 += y2Add

			sum += math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
		}
	}

	fmt.Println(int(sum))
}
