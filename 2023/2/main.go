package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SetPair struct {
	Color  string
	Number int
}

type Set struct {
	SetPairs []SetPair
}

type Game struct {
	Number int
	Sets   []Set
}

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

	var games []Game

	for _, gameLine := range textSlice {
		gameSplit := strings.Split(gameLine, ":")

		gameString := strings.Fields(gameSplit[0])[1]
		gameInt, _ := strconv.Atoi(gameString)

		sets := strings.Split(gameSplit[1], ";")

		var usableSets []Set

		for _, set := range sets {
			colors := strings.Split(set, ",")

			var gameSet Set
			var setPairs []SetPair
			for _, color := range colors {
				elements := strings.Fields(color)

				if elements[1] == "blue" {

					drawInt, _ := strconv.Atoi(elements[0])

					setPair := SetPair{
						Color:  "blue",
						Number: drawInt,
					}

					setPairs = append(setPairs, setPair)
				}

				if elements[1] == "green" {
					drawInt, _ := strconv.Atoi(elements[0])

					setPair := SetPair{
						Color:  "green",
						Number: drawInt,
					}

					setPairs = append(setPairs, setPair)
				}

				if elements[1] == "red" {
					drawInt, _ := strconv.Atoi(elements[0])

					setPair := SetPair{
						Color:  "red",
						Number: drawInt,
					}

					setPairs = append(setPairs, setPair)
				}

				gameSet = Set{
					SetPairs: setPairs,
				}
				usableSets = append(usableSets, gameSet)
			}

		}
		usableGame := Game{
			Number: gameInt,
			Sets:   usableSets,
		}
		games = append(games, usableGame)
	}

	var matchingGames []int

	for _, game := range games {
		maxRed, maxBlue, maxGreen := 0, 0, 0
		for _, set := range game.Sets {
			for _, pair := range set.SetPairs {
				if pair.Color == "red" {
					if pair.Number > maxRed {
						maxRed = pair.Number
					}

				}

				if pair.Color == "blue" {
					if pair.Number > maxBlue {
						maxBlue = pair.Number
					}
				}

				if pair.Color == "green" {
					if pair.Number > maxGreen {
						maxGreen = pair.Number
					}
				}
			}
		}

		matchingGames = append(matchingGames, maxRed*maxBlue*maxGreen)
	}

	sum := 0
	for _, num := range matchingGames {
		sum += num
	}

	fmt.Println(matchingGames)
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
