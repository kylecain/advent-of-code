package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    leftList := []int{}
    rightList := []int{}

    for scanner.Scan() {
        line := scanner.Text()
        locationId := strings.Fields(line)

        leftId, err := strconv.Atoi(locationId[0])
        if err != nil {
            fmt.Println(err)
            return
        }

        rightId, _ := strconv.Atoi(locationId[1])
        if err != nil {
            fmt.Println(err)
            return
        }

        leftList = append(leftList, leftId)
        rightList = append(rightList, rightId)
    }

    rightMap := make(map[int]int)

    for _, rightId := range(rightList) {
        _, exists := rightMap[rightId]
        if exists {
            rightMap[rightId]++
        } else {
            rightMap[rightId] = 1
        }
    }

    rightMap[0] = 0

    sum := 0

    for _, leftId := range(leftList) {
        sum += leftId * rightMap[leftId]
    }

    fmt.Println(sum)
}
