package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"regexp"
)

type Robot struct {
    p Point
    moves []string
}

type Point struct {
    i int
    j int
}

var up Point = Point{-1, 0}
var down Point = Point{1, 0}
var left Point = Point{0, -1}
var right Point = Point{0, 1}

var directionMap = map[string]Point {
    "^": up,
    "v": down,
    "<": left,
    ">": right,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

    // partOne(file)
    partTwo(file)
}

func partOne(file io.Reader) {
	scanner := bufio.NewScanner(file)
    var robot Robot
    var grid [][]string
    var i int
    isReadingGrid := true

	for scanner.Scan() {
		line := scanner.Text()

        if len(line) == 0 {
            isReadingGrid = false
            continue
        }

        if isReadingGrid {
            exp := regexp.MustCompile(`@`)
            robotIndex := exp.FindStringIndex(line)
            if len(robotIndex) > 0 {
                robot = Robot{Point{i, robotIndex[0]}, []string{}}
            }
            
            temp := []string{}
            for _, r := range line {
                temp = append(temp, string(r))
            }
            grid = append(grid, temp)
        } else {
            for _, r := range line {
                robot.moves = append(robot.moves, string(r))
            }

        }

        i++
	}
    for len(robot.moves) > 0 {
        robot.move(grid, 1)
    }
    
    fmt.Println(getSum(grid))
}

func partTwo(file io.Reader) {
	scanner := bufio.NewScanner(file)
    var robot Robot
    var grid [][]string
    var i int
    isReadingGrid := true

	for scanner.Scan() {
		line := scanner.Text()

        if len(line) == 0 {
            isReadingGrid = false
            continue
        }

        if isReadingGrid {
            temp := []string{}
            for _, r := range line {
                if string(r) == "#" {
                    temp = append(temp, string(r))
                    temp = append(temp, string(r))
                } else if string(r) == "O" {
                    temp = append(temp, "[")
                    temp = append(temp, "]")
                } else if string(r) == "@" {
                    temp = append(temp, string(r))
                    temp = append(temp, ".")
                } else {
                    temp = append(temp, string(r))
                    temp = append(temp, string(r))
                }
            }
            grid = append(grid, temp)
        } else {
            for _, r := range line {
                robot.moves = append(robot.moves, string(r))
            }

        }

        i++
	}

    for i, v := range grid {
        for j, v1 := range v {
            if v1 == "@" {
                robot.p = Point{i, j}
            }
        }
    }

    printGrid(grid)
    for len(robot.moves) > 0 {
        robot.move(grid, 2)
        printGrid(grid)
    }

    fmt.Println(getSum(grid))
}

func getSum(grid [][]string) int {
    sum := 0 
    for i, v := range grid {
        for j, v1 := range v {
            if v1 == "O" {
                sum += 100 * i + j
            }
        }
    }

    return sum
}

func (r *Robot) move(grid [][]string, part int) {
    move := r.moves[0]
    r.moves = r.moves[1:]
    p := directionMap[move]
    fmt.Println("move: ")

    if part == 1 {
        if r.moveInDirection(grid, r.p.i, r.p.j, p) {
            r.p = Point{r.p.i + p.i, r.p.j + p.j}
        }
    } else {
        if r.moveInDirection2(grid, r.p, p) {
            r.p = Point{r.p.i + p.i, r.p.j + p.j}
        }
    }
}

func (r *Robot) moveInDirection(grid [][]string, i int, j int, p Point) bool {
    i1, j1 := i + p.i, j + p.j
    char := grid[i1][j1]

    if char == "#" {
        return false
    } else if char == "O" {
        if !r.moveInDirection(grid, i1, j1, p) {
            return false
        }
    } 

    grid[i1][j1] = grid[i][j]
    grid[i][j] = "."

    return true
}

func (r *Robot) moveInDirection2(grid [][]string, current Point, direction Point) bool {
    i1, j1 := current.i + direction.i, current.j + direction.j
    characterAhead := grid[i1][j1]
    pointsAhead := []Point{{i1, j1}}

    if characterAhead == "#" {
        return false
    } else if characterAhead == "." {
        grid[i1][j1], grid[current.i][current.j] = grid[current.i][current.j], grid[i1][j1]
        return true
    } else if characterAhead == "]" {
        pointsAhead = append(pointsAhead, Point{i1, j1 + left.j})
    } else if characterAhead == "[" {
        pointsAhead = append(pointsAhead, Point{i1, j1 + right.j})
    }

    if direction == up || direction == down {
        for _, point := range pointsAhead {
            if r.moveInDirection2(grid, point, direction) {
                fmt.Println("swapping:",grid[i1][j1], grid[current.i][current.j])
                grid[i1][j1], grid[current.i][current.j] = grid[current.i][current.j], grid[i1][j1]
            }
        }
    } else if direction == left || direction == right {
        if r.moveInDirection2(grid, pointsAhead[0], direction) {
            grid[i1][j1], grid[current.i][current.j] = grid[current.i][current.j], grid[i1][j1]
        }
    }

    return true
}



func printGrid(grid [][]string) {
    for _, v := range grid {
        fmt.Println(v)
    }
}

