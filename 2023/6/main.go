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

	var timesAndDistances [][]int
	numsArrayIndex := -1

	for _, str := range numsArray {
		if str == ":" {
			numsArrayIndex++
			timesAndDistances = append(timesAndDistances, []int{})
			continue
		}
		num, _ := strconv.Atoi(str)
		timesAndDistances[numsArrayIndex] = append(timesAndDistances[numsArrayIndex], num)
	}

	times := timesAndDistances[0]
	distances := timesAndDistances[1]

	timeString := ""
	distanceString := ""
	for i := 0; i < len(times); i++ {
		timeString += strconv.Itoa(times[i])
		distanceString += strconv.Itoa(distances[i])
	}
	timeInt, _ := strconv.Atoi(timeString)
	distanceInt, _ := strconv.Atoi(distanceString)
	times = []int{timeInt}
	distances = []int{distanceInt}

	var winsPerGame []int

	for i := 0; i < len(times); i++ {
		wins := 0
		for j := 0; j <= times[i]; j++ {
			timeToTravel := times[i] - j
			distanceTraveled := timeToTravel * j

			if distanceTraveled > distances[i] {
				wins++
			}
		}
		winsPerGame = append(winsPerGame, wins)
	}

	p := 1
	for _, wins := range winsPerGame {
		p = p * wins
	}

	fmt.Println(p)
}
