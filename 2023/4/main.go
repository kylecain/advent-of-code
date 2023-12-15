package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

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
	
	gameNumberToMatches := make(map[int]int)

	sum := 0
	for gameIndex, lineOfText := range textSlice {
		splitText := strings.Split(lineOfText, ":")
		cards := strings.Split(splitText[1], "|")

		var strippedCards [][]string
		for _, card := range cards {
			strippedCards = append(strippedCards, strings.Fields(card))
		}

		var setOfCardsAsInts [][]int
		for _, setOfCards := range strippedCards {
			var arrayOfCards []int
			for _, cardStr := range setOfCards {
				cardInt, _ := strconv.Atoi(string(cardStr))
				arrayOfCards = append(arrayOfCards, cardInt)
			}
			setOfCardsAsInts = append(setOfCardsAsInts, arrayOfCards)
		}

		var matchesMap = map[int]int{}

		for _, setOfCards := range setOfCardsAsInts {
			for _, card := range setOfCards {
				value, exists := matchesMap[card]
				if exists {
					matchesMap[card] = value+1
				} else {
					matchesMap[card] = 1
				}
			}
		}

		var winningNumbers = []int{}
		for key, value := range matchesMap {
			if value > 1 {
				winningNumbers = append(winningNumbers, key)
			}
		}
		
		if len(winningNumbers) > 0 {
			exp := len(winningNumbers)-1
			sum += intPow(2, exp)
			gameNumberToMatches[gameIndex] = len(winningNumbers)
		} else {
			gameNumberToMatches[gameIndex] = 0
		}
	}

	mapLength := len(gameNumberToMatches)
	numOfCardsArray := make([]int, mapLength)
	var numOfMatchesArray []int
	for i := range numOfCardsArray {
		numOfCardsArray[i] = 1
	}
	
	for i:=0;i<mapLength;i++ {
		numOfMatchesArray = append(numOfMatchesArray, gameNumberToMatches[i])
	}

	for index, value := range numOfMatchesArray {
		for i:=1;i<=value;i++ {
			numOfCardsArray[index+i] = numOfCardsArray[index] + numOfCardsArray[index+i]
		}
	}

	numOfCardsSum := 0
	for _, value := range numOfCardsArray {
		numOfCardsSum += value
	}

	fmt.Println(gameNumberToMatches)
	fmt.Println(numOfCardsArray)
	fmt.Println(numOfCardsSum)

}

func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}


