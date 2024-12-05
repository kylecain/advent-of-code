package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	sum := 0
	isAdding := true

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\([0-9]{1,3},[0-9]{1,3}\))`)

		matches := r.FindAllString(line, -1)

		for _, match := range matches {
			if match == "do()" {
				isAdding = true
				continue
			} else if match == "don't()" {
				isAdding = false
				continue
			}

			if !isAdding {
				continue
			}

			r1 := regexp.MustCompile(`[0-9]{1,3}`)
			m1 := r1.FindAllString(match, 2)

			x, _ := strconv.Atoi(m1[0])
			y, _ := strconv.Atoi(m1[1])

			fmt.Println(isAdding, x, y)

			sum += x * y
		}
	}
	fmt.Println(sum)
}
