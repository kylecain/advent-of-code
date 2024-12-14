package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
    p Point
    vi int
    vj int
}

type Point struct {
    i int
    j int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
    robots := []*Robot{}

	for scanner.Scan() {
		line := scanner.Text()

        exp := regexp.MustCompile(`-?\d+`)
        nums := exp.FindAllString(line, -1)

        pi, _ := strconv.Atoi(nums[1])
        pj, _ := strconv.Atoi(nums[0])
        vi, _ := strconv.Atoi(nums[3])
        vj, _ := strconv.Atoi(nums[2])

        p := Point{pi, pj}

        robots = append(robots, &Robot{p, vi, vj})
	}

    // partOne(robots)
    partTwo(robots)
}

func partOne(robots []*Robot) {
    var grid [103][101]int

    for _, robot := range robots {
        grid[robot.p.i][robot.p.j]++
    }

    for i := 1; i <= 100; i++ {
        for _, robot := range robots {
            robot.move(&grid)
        }
    }

    tl, tr, bl, br := 0, 0, 0 ,0
    iMid, jMid := len(grid) / 2, len(grid[0]) / 2

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if i < iMid && j < jMid {
                tl += grid[i][j]
            }

            if i > iMid && j < jMid {
                bl += grid[i][j]
            }

            if i < iMid && j > jMid {
                tr += grid[i][j]
            }

            if i > iMid && j > jMid {
                br += grid[i][j]
            }
        }
    }

    total := tl * tr * bl * br
    fmt.Println(total)
}

func partTwo(robots []*Robot) {
    var grid [103][101]int

    file, _ := os.Create("robots.txt")
    defer file.Close()

    writer := bufio.NewWriter(file)

    for _, robot := range robots {
        grid[robot.p.i][robot.p.j]++
    }

    for i := 1; i <= 100000; i++ {
        for _, robot := range robots {
            robot.move(&grid)
        }

        iteration := "iteration: " + strconv.Itoa(i) + "\n"
        writer.WriteString(iteration)
        var vStr string
        for _, v := range grid {
            for _, v1 := range v {
                if v1 == 0 {
                    vStr += "."
                } else {
                    vStr += strconv.Itoa(v1)
                }
            }
            vStr += "\n"
        }
        vStr += "\n"
        writer.WriteString(vStr)
    }
}

func (r *Robot) move(grid *[103][101]int) {
    grid[r.p.i][r.p.j]--
    r.p.i += r.vi
    r.p.j += r.vj

    ilen := len(grid)
    jlen := len(grid[0])

    if r.p.i > ilen -1 {
        r.p.i -= ilen
    }

    if r.p.j > jlen -1 {
        r.p.j -= jlen
    }

    if r.p.i < 0 {
        r.p.i += ilen
    }

    if r.p.j < 0 {
        r.p.j += jlen
    }

    grid[r.p.i][r.p.j]++
}

func printGrid(grid [103][101]int) {
    for _, v := range grid {
        fmt.Println(v)
    }
}
