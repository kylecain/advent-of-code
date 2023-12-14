package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numsMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

var numsSliceForwards = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var numsSliceBackwards = []string{
	"eno",
	"owt",
	"eerht",
	"ruof",
	"evif",
	"xis",
	"neves",
	"thgie",
	"enin",
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

	var backwardsSlice, addedNumbersSlice []string

	for _, text := range textSlice {
		backwardsSlice = append(backwardsSlice, reverseString(text))
	}

	forwardsNumbers := findFirstNumbers(textSlice, numsSliceForwards)
	backwardsNumbers := findFirstNumbers(backwardsSlice, numsSliceBackwards)

	for i := 0; i < len(forwardsNumbers); i++ {
		addedNumbersSlice = append(addedNumbersSlice, forwardsNumbers[i]+backwardsNumbers[i])
	}

	result := 0
	for _, value := range addedNumbersSlice {
		num, _ := strconv.Atoi(value)
		result += num
	}
	fmt.Println(result)
}

func findFirstNumbers(str []string, keys []string) []string {
	var numbers []string
	for _, text := range str {
		var firstNum string

		for i := 0; i < len(text); i++ {
			str := string(text[i])

			if regexp.MustCompile(`^\d$`).MatchString(str) {
				firstNum = str
				break
			} else {
				for _, numberString := range keys {
					if i+len(numberString) > len(text) {
						continue
					}
					if text[i:i+len(numberString)] == numberString {
						firstNum = numsMap[numberString]
						break
					}
				}
				if len(firstNum) != 0 {
					break
				}
			}
		}

		numbers = append(numbers, firstNum)
	}

	return numbers
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
