package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardValues map[string]int

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Hand  string
	Bid   int
	Score int
}

func main() {
	cardValues = map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hands []Hand

	for scanner.Scan() {
		handMap := make(map[rune]int)
		line := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(line[1])

		for _, r := range line[0] {
			handMap[r]++
		}

		highestValue, secondHighestValue, jokerCount := 0, 0, 0
		for key, value := range handMap {
			if key == 74 {
				jokerCount = value
			} else {
				if value >= highestValue {
					secondHighestValue = highestValue
					highestValue = value
				} else if value >= secondHighestValue {
					secondHighestValue = value
				}
			}
		}

		highestValue = highestValue + jokerCount

		score := 0
		if highestValue == 5 {
			score = FiveOfAKind
		} else if highestValue == 4 {
			score = FourOfAKind
		} else if highestValue == 3 && secondHighestValue == 2 {
			score = FullHouse
		} else if highestValue == 3 {
			score = ThreeOfAKind
		} else if highestValue == 2 && secondHighestValue == 2 {
			score = TwoPair
		} else if highestValue == 2 {
			score = OnePair
		} else {
			score = HighCard
		}

		hand := Hand{
			Hand:  line[0],
			Bid:   bid,
			Score: score,
		}

		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Score != hands[j].Score {
			return hands[i].Score < hands[j].Score
		}

		for k := 0; k < 5; k++ {
			handI := cardValues[string(hands[i].Hand[k])]
			handJ := cardValues[string(hands[j].Hand[k])]

			if handI != handJ {
				return handI < handJ
			}
		}

		return false
	})

	s := 0
	for i := 0; i < len(hands); i++ {
		s += hands[i].Bid * (i + 1)
	}

	fmt.Println(s)
}
