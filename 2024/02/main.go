package main

import (
	"bufio"
	"errors"
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

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)

		var levelSlice []int
		for _, level := range levels {
			levelInt, _ := strconv.Atoi(level)
			levelSlice = append(levelSlice, levelInt)
		}

		fmt.Println(levelSlice)

		isSafe, err := CheckNums(levelSlice)

		if err != nil {
			fmt.Println(err)
			return
		}

		if isSafe {
			safeCount++
			continue
		}

		for i := 0; i < len(levelSlice); i++ {
			modifiedLevels := RemoveIndex(levelSlice, i)

			isSafe, err := CheckNums(modifiedLevels)
			if err != nil {
				fmt.Println(err)
				return
			}

			if isSafe {
				safeCount++
				break
			}
		}

	}

	fmt.Println(safeCount)
}

func CheckNums(levels []int) (bool, error) {
	direction := ""
	previousLevel := levels[0]

	for i := 1; i < len(levels); i++ {
		currentLevel := levels[i]

		if currentLevel > previousLevel {
			if direction == "decending" {
				return false, nil
			}

			direction = "accending"

		} else if currentLevel < previousLevel {
			if direction == "accending" {
				return false, nil
			}

			direction = "decending"
		} else {
			return false, nil
		}

		levelDifference := currentLevel - previousLevel
		absLevelDifference := max(levelDifference, -levelDifference)

		if absLevelDifference > 3 {
			return false, nil
		}

		if i == len(levels)-1 {
			return true, nil
		} else {
			previousLevel = currentLevel
		}
	}

	return false, errors.New("unexpected error")
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
