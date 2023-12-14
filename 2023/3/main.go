package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Star struct {
	X int
	Y int
}

var starsMap = map[Star][]int{}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var textSlice []string

	for scanner.Scan() {
		textSlice = append(textSlice, scanner.Text())
	}

	totalRows := len(textSlice)
	digitRegex := regexp.MustCompile(`\d+`)
	symbolRegex := regexp.MustCompile(`\*`)

	sum := 0

	for rowIndex, str := range textSlice {
		numIndexes := digitRegex.FindAllStringIndex(str, -1)

		for _, numIndex := range numIndexes {
			numAsString := string(str[numIndex[0]:numIndex[1]])
			var startRange, endRange int
			if numIndex[0]-1 < 0 {
				startRange = numIndex[0]
			} else {
				startRange = numIndex[0] - 1
			}
			if numIndex[1]+1 > len(str) {
				endRange = numIndex[1]
			} else {
				endRange = numIndex[1] + 1
			}

			// looks at row above
			if rowIndex-1 > 0 {
				stringInQuestion := textSlice[rowIndex-1][startRange:endRange]

				starIndexes := symbolRegex.FindAllStringIndex(stringInQuestion, -1)

				for _, starIndex := range starIndexes {
					numAsInt, _ := strconv.Atoi(numAsString)
					addStar(rowIndex-1, startRange+starIndex[0], numAsInt)
				}
			}

			// looks at row below
			if rowIndex+1 < totalRows {
				stringInQuestion := textSlice[rowIndex+1][startRange:endRange]
				starIndexes := symbolRegex.FindAllStringIndex(stringInQuestion, -1)

				for _, starIndex := range starIndexes {
					numAsInt, _ := strconv.Atoi(numAsString)
					addStar(rowIndex+1, startRange+starIndex[0], numAsInt)
				}
			}

			stringInQuestion := str[startRange:endRange]
			starIndexes := symbolRegex.FindAllStringIndex(stringInQuestion, -1)

			for _, starIndex := range starIndexes {
				numAsInt, _ := strconv.Atoi(numAsString)
				addStar(rowIndex, startRange+starIndex[0], numAsInt)
			}
		}
	}

	for _, value := range starsMap {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}

	fmt.Println(sum)
}

func addStar(x int, y int, number int) {
	star := Star{X: x, Y: y}
	value, exists := starsMap[star]

	if exists {
		starsMap[star] = append(value, number)
	} else {
		starsMap[star] = []int{number}
	}
}
