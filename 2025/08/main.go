package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

type Point struct {
	X           int
	Y           int
	Z           int
	Connections []*Point
}

func (p1 *Point) DistanceTo(p2 Point) float64 {
	dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z
	xx, yy, zz := dx*dx, dy*dy, dz*dz
	sum := float64(xx + yy + zz)
	return math.Sqrt(sum)
}

func (p1 *Point) AddConnection(p2 *Point) {
	p1.Connections = append(p1.Connections, p2)
}

func (p1 *Point) IsConnected(p2 *Point) bool {
	return slices.Contains(p1.Connections, p2)
}

func (p1 *Point) GetCircuitSize() int {
	visited := make(map[*Point]bool)
	dfs(p1, visited)
	return (len(visited))
}

func dfs(p *Point, visited map[*Point]bool) {
	if visited[p] {
		return
	}
	visited[p] = true
	for _, c := range p.Connections {
		dfs(c, visited)
	}
}

func ConnectPoints(p1, p2 *Point) {
	p1.AddConnection(p2)
	p2.AddConnection(p1)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var points []*Point
	for scanner.Scan() {
		x, y, z := 0, 0, 0
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		points = append(points, &Point{x, y, z, []*Point{}})
	}

	fmt.Println(part1(points))
	fmt.Println(part2(points))
}

func part1(points []*Point) int {
	numPairs := 1000
	numCircuits := 3

	for range numPairs {
		min := math.MaxFloat64
		var p1, p2 *Point
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				if points[i].IsConnected(points[j]) {
					continue
				}

				d := points[i].DistanceTo(*points[j])
				if d < min {
					min = d
					p1, p2 = points[i], points[j]
				}

			}
		}

		ConnectPoints(p1, p2)
	}

	var lengths []int
	visited := make(map[*Point]bool)
	for _, p := range points {
		if visited[p] {
			continue
		}
		dfs(p, visited)
		lengths = append(lengths, p.GetCircuitSize())
	}

	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] > lengths[j]
	})

	product := 1
	for _, v := range lengths[:numCircuits] {
		product *= v
	}

	return product
}

func part2(points []*Point) int {
	var pp1, pp2 *Point

	for {
		min := math.MaxFloat64
		var p1, p2 *Point
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				if points[i].IsConnected(points[j]) {
					continue
				}

				d := points[i].DistanceTo(*points[j])
				if d < min {
					min = d
					p1, p2 = points[i], points[j]
				}

			}
		}

		ConnectPoints(p1, p2)
		pp1, pp2 = p1, p2

		shouldBreak := true
		for _, p := range points {
			if len(p.Connections) == 0 {
				shouldBreak = false
			}
		}
		if shouldBreak {
			break
		}
	}

	return pp1.X * pp2.X
}
