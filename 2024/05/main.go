package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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

	rules := [][]string{}
	updates := [][]string{}

	readingRules := true

	for scanner.Scan() {
		line := scanner.Text()

		exp := regexp.MustCompile("[0-9]{2}")
		numbers := exp.FindAllString(line, -1)

		if len(numbers) == 0 {
			readingRules = false
			continue
		}

		if readingRules {
			rules = append(rules, numbers)
		} else {
			updates = append(updates, numbers)
		}

	}

	// partOne(rules, updates)
	partTwo(rules, updates)
}

func partOne(rules, updates [][]string) {
	ruleMap := make(map[string][]string)

	sum := 0

	for _, rule := range rules {
		ruleKey := rule[1]
		ruleValue := rule[0]

		ruleMap[ruleKey] = append(ruleMap[ruleKey], ruleValue)
	}

	for _, update := range updates {
		isValid := true

		for i, page := range update {
			invalidPages := ruleMap[page]

			fmt.Println(page, invalidPages)

			for j := i + 1; j < len(update); j++ {
				for _, invalidPage := range invalidPages {
					if update[j] == invalidPage {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}

			if !isValid {
				break
			}
		}

		if isValid {
			middlePage, _ := strconv.Atoi(update[len(update)/2])
			sum += middlePage
		}
	}

	fmt.Println(sum)
}

func partTwo(rules, updates [][]string) {
	ruleMap := make(map[string][]string)
	invalidUpdates := [][]string{}

	for _, rule := range rules {
		ruleKey := rule[1]
		ruleValue := rule[0]

		ruleMap[ruleKey] = append(ruleMap[ruleKey], ruleValue)
	}

	for _, update := range updates {
		isValid := true

		for i, page := range update {
			invalidPages := ruleMap[page]

			for j := i + 1; j < len(update); j++ {
				for _, invalidPage := range invalidPages {
					if update[j] == invalidPage {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}

			if !isValid {
				break
			}
		}

		if !isValid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	sum := 0

	for _, update := range invalidUpdates {
		slices.SortFunc(update, func(a, b string) int {
			if slices.Contains(ruleMap[a], b) {
				return 1
			} else {
				return -1
			}
		})

		middlePage, _ := strconv.Atoi(update[len(update)/2])
		sum += middlePage
	}

	fmt.Println(sum)
}
