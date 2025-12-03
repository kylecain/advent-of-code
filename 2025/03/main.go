package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var banks []string
	for scanner.Scan() {
		line := scanner.Text()

		banks = append(banks, line)
	}

	part1(banks)
	part2(banks)
}

func part1(banks []string) {
	sum := 0
	for _, bank := range banks {
		var maxILeft int
		var maxBLeft, maxBRight rune

		for i, b := range bank[:len(bank)-1] {
			if b > maxBLeft {
				maxBLeft = b
				maxILeft = i
			}
		}

		for _, b := range bank[maxILeft+1:] {
			if b > maxBRight {
				maxBRight = b
			}
		}

		c := string(maxBLeft) + string(maxBRight)
		i, _ := strconv.Atoi(c)
		sum += i
	}
	fmt.Println(sum)
}

func part2(banks []string) {
	sum := 0
	for _, bank := range banks {
		var batteries []rune
		leftBound, rightNeeded := 0, 12

		for rightNeeded > 0 {
			var max rune
			localLeftBound := 0
			rightBound := len(bank) - rightNeeded + 1

			for k, b := range bank[leftBound:rightBound] {
				if b > max {
					max = b
					localLeftBound = k
				}
			}

			leftBound = leftBound + localLeftBound + 1
			batteries = append(batteries, max)
			rightNeeded--
		}

		intValue, _ := strconv.Atoi(string(batteries))
		sum += intValue
	}
	fmt.Println(sum)
}
