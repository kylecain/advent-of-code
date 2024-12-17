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
    fmt.Println("move: ", move)

    if part == 1 {
        if r.moveInDirection(grid, r.p.i, r.p.j, p) {
            r.p = Point{r.p.i + p.i, r.p.j + p.j}
        }
    } else {
        pointsAhead := []Point{Point{r.p.i + p. i, r.p.j + p.j}}
        pointsBehind := []Point{Point{r.p.i, r.p.j}}
        charAhead := grid[r.p.i + p.i][r.p.j + p.j]
        if charAhead == "[" && (p == up || p == down){
            pointsAhead = append(pointsAhead, Point{pointsAhead[0].i, pointsAhead[0].j + right.j})
            pointsBehind= append(pointsBehind, Point{pointsBehind[0].i, pointsBehind[0].j + right.j})
        }
        if charAhead == "]" && (p == up || p == down){
            pointsAhead = append(pointsAhead, Point{pointsAhead[0].i, pointsAhead[0].j + left.j})
            pointsBehind= append(pointsBehind, Point{pointsBehind[0].i, pointsBehind[0].j + left.j})
        }
        if r.moveInDirection2(grid, p, pointsAhead, pointsBehind) {
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

func (r *Robot) moveInDirection2(grid [][]string, direction Point, pointsAhead []Point, pointsBehind []Point) bool {
    if len(pointsAhead) > 1 {
        if grid[pointsAhead[0].i][pointsAhead[0].j] == "#" || grid[pointsAhead[1].i][pointsAhead[1].j] == "#" {
            return false
        }
    } else {
        if grid[pointsAhead[0].i][pointsAhead[0].j] == "#" {
            return false
        }
    }

    if len(pointsAhead) > 1 {
        if grid[pointsAhead[0].i][pointsAhead[0].j] == "." && grid[pointsAhead[1].i][pointsAhead[1].j] == "." {
            grid[pointsAhead[0].i][pointsAhead[0].j], grid[pointsBehind[0].i][pointsBehind[0].j] = grid[pointsBehind[0].i][pointsBehind[0].j], grid[pointsAhead[0].i][pointsAhead[0].j]
            grid[pointsAhead[1].i][pointsAhead[1].j], grid[pointsBehind[1].i][pointsBehind[1].j] = grid[pointsBehind[1].i][pointsBehind[1].j], grid[pointsAhead[1].i][pointsAhead[1].j]
            printGrid(grid)
            return true
        }
    } else {
        if grid[pointsAhead[0].i][pointsAhead[0].j] == "." {
            grid[pointsAhead[0].i][pointsAhead[0].j], grid[pointsBehind[0].i][pointsBehind[0].j] = grid[pointsBehind[0].i][pointsBehind[0].j], grid[pointsAhead[0].i][pointsAhead[0].j]
            return true
        }
    }


    if direction == up || direction == down {
        newPointsAhead := []Point{Point{pointsAhead[0].i + direction.i, pointsAhead[0].j + direction.j},Point{pointsAhead[1].i + direction.i, pointsAhead[1].j + direction.j}}
        if r.moveInDirection2(grid, direction, newPointsAhead, pointsAhead) && r.moveInDirection2(grid, direction, newPointsAhead, pointsAhead) {
            fmt.Println("swapping:",grid[pointsAhead[0].i][pointsAhead[0].j], grid[pointsBehind[0].i][pointsBehind[0].j])
            fmt.Println("swapping:",grid[pointsAhead[1].i][pointsAhead[1].j], grid[pointsBehind[1].i][pointsBehind[1].j])
            fmt.Println()

            grid[pointsAhead[0].i][pointsAhead[0].j], grid[pointsBehind[0].i][pointsBehind[0].j] = grid[pointsBehind[0].i][pointsBehind[0].j], grid[pointsAhead[0].i][pointsAhead[0].j]
            grid[pointsAhead[1].i][pointsAhead[1].j], grid[pointsBehind[1].i][pointsBehind[1].j] = grid[pointsBehind[1].i][pointsBehind[1].j], grid[pointsAhead[1].i][pointsAhead[1].j]
            printGrid(grid)
        }
    } else if direction == left || direction == right {
        newPointsAhead := []Point{Point{pointsAhead[0].i + direction.i, pointsAhead[0].j + direction.j}}
        if r.moveInDirection2(grid, direction, newPointsAhead, pointsAhead) {
            grid[pointsAhead[0].i][pointsAhead[0].j], grid[pointsBehind[0].i][pointsBehind[0].j] = grid[pointsBehind[0].i][pointsBehind[0].j], grid[pointsAhead[0].i][pointsAhead[0].j]
        }
    }

    return true
}



func printGrid(grid [][]string) {
    for _, v := range grid {
        fmt.Println(v)
    }
}

