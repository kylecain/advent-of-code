package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("cant open file")
	}
	fileStr := strings.TrimSpace(string(file))
	ranges := strings.Split(fileStr, ",")

	part1(ranges)
	part2(ranges)
}

func part1(ranges []string) {
	sum := 0
	for _, v := range ranges {
		ids := strings.Split(v, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			panic("cant read start")
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			panic("can't read end")
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 == 1 {
				continue
			}
			if s[:len(s)/2] == s[len(s)/2:] {
				sum += i
			}
		}
	}

	fmt.Println("sum is:", sum)

}

func part2(ranges []string) {
	sum := 0
	for _, v := range ranges {
		ids := strings.Split(v, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			panic("cant read start")
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			panic("can't read end")
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			size := len(s)

			for j := 1; j <= size/2; j++ {
				p := s[:j]

				isValid := true

				for k := j; k < size; k += j {

					if k+j > size {
						isValid = false
						break
					}
					if s[k:k+j] != p {
						isValid = false
						break
					}
				}

				if isValid {
					sum += i
					break
				}
			}

		}
	}

	fmt.Println("sum is:", sum)
}
