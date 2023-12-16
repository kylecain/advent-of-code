package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	characterRegex := regexp.MustCompile("[a-zA-Z]|-")
	filteredData := characterRegex.ReplaceAll(data, []byte(""))
	numsArray := strings.Fields(string(filteredData))

	var seedsAndMaps [][]int
	numsArrayIndex := -1

	for _, str := range numsArray {
		if str == ":" {
			numsArrayIndex++
			seedsAndMaps = append(seedsAndMaps, []int{})
			continue
		}
		num, _ := strconv.Atoi(str)
		seedsAndMaps[numsArrayIndex] = append(seedsAndMaps[numsArrayIndex], num)
	}

	seeds := seedsAndMaps[0]
	k := 0
	for {
		seedFound := false
		currentValue := k

		for i := len(seedsAndMaps) - 1; i >= 1; i-- {
			currentMap := seedsAndMaps[i]

			for j := 0; j < len(currentMap); j = j + 3 {
				lowerBound := currentMap[j]
				upperBound := currentMap[j] + currentMap[j+2] - 1
				if currentValue >= lowerBound && currentValue <= upperBound {
					currentValue = currentValue - currentMap[j] + currentMap[j+1]
					break
				}
			}
		}

		for m := 0; m < len(seeds); m = m + 2 {
			if currentValue >= seeds[m] && currentValue <= seeds[m]+seeds[m+1] {
				seedFound = true
				break
			}
		}

		if seedFound {
			break
		} else {
			k++
		}
	}

	fmt.Println(k)
}
