package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	Row, Col int
	Active   bool
}

type Region struct {
	Rows, Cols int
	ShapeCount []int
}

func NewRegion(definition string) *Region {
	parts := strings.Split(definition, ":")
	dimensions := parts[0]
	activeStr := strings.Fields(parts[1])

	var rows, cols int
	fmt.Sscanf(dimensions, "%dx%d", &rows, &cols)

	active := make([]int, len(activeStr))
	for i, s := range activeStr {
		intValue, _ := strconv.Atoi(s)
		active[i] = intValue
	}
	return &Region{rows, cols, active}
}

func (r *Region) Area() int {
	return r.Rows * r.Cols
}

type Shape struct {
	Points [][]Point
}

func NewShape(shapeStr []string) *Shape {
	points := make([][]Point, len(shapeStr))
	for row := range points {
		points[row] = make([]Point, len(shapeStr[0]))
	}
	for row, s := range shapeStr {
		for col, r := range s {
			active := false
			if r == '#' {
				active = true
			}

			points[row][col] = Point{row, col, active}
		}
	}
	return &Shape{points}
}

func (s *Shape) Height() int {
	return len(s.Points)
}

func (s *Shape) Width() int {
	return len(s.Points[0])
}

func (s *Shape) Area() int {
	a := 0
	for _, row := range s.Points {
		for _, p := range row {
			if p.Active {
				a++
			}

		}
	}
	return a
}

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.Split(string(input), "\n\n")

	shapes := make(map[int]*Shape)
	for i, s := range data[:len(data)-1] {
		shapes[i] = NewShape(strings.Split(s, "\n")[1:])
	}

	rs := strings.Split(data[len(data)-1], "\n")
	regions := make(map[int]*Region)
	for i, s := range rs[:len(rs)-1] {
		regions[i] = NewRegion(s)

	}

	fmt.Println(part1(shapes, regions))
}

func part1(shapes map[int]*Shape, regions map[int]*Region) int {
	sum := 0
	for _, r := range regions {
		regionArea := r.Area()

		shapeArea := 0
		for i, v := range r.ShapeCount {
			shapeArea += shapes[i].Area() * v
		}

		if shapeArea < regionArea {
			sum++
		}
	}
	return sum
}
