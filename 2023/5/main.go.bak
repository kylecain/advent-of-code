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

	for index, seed := range seeds {
		value := seed
		for i := 1; i < len(seedsAndMaps); i++ {
			currentMap := seedsAndMaps[i]

			// check if between sources
			for j := 1; j < len(currentMap); j = j + 3 {
				lowerBound := currentMap[j]
				upperBound := currentMap[j] + currentMap[j+1] - 1
				if value >= lowerBound && value <= upperBound {
					// if source matches look at desination
					value = value - currentMap[j] + currentMap[j-1]
					break
				}
			}
		}
		seeds[index] = value
	}

	lowestNum := seeds[0]
	for i := 1; i < len(seeds); i++ {
		if seeds[i] < lowestNum {
			lowestNum = seeds[i]
		}
	}

	fmt.Println(lowestNum)
}
