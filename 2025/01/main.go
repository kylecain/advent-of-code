package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := 50
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		d := string(line[0])
		n := line[1:]
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
			return
		}

		if d == "L" {
			pos -= i
		} else {
			pos += i
		}

		if pos%100 == 0 {
			sum++
		}
	}

	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := 50
	sum := 0
	fmt.Println(pos)
	for scanner.Scan() {
		line := scanner.Text()
		d := string(line[0])
		n := line[1:]
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
			return
		}

		if d == "L" {
			if i >= pos && pos != 0 {
				sum++
			}

			pos -= i

			q, r := pos/100, pos%100

			if r < 0 {
				pos = r + 100
			} else {
				pos = r
			}

			sum += int(math.Abs(float64(q)))
		} else {
			pos += i
			q, r := pos/100, pos%100

			pos = r

			sum += q
		}

		fmt.Println(line, pos, sum)
	}

	fmt.Println(sum)
}
