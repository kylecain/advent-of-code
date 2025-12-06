package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var problems [][]string

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		problems = append(problems, fields)
	}
	numbers := problems[:len(problems)-1]
	signs := problems[len(problems)-1:][0]
	var answers []int
	for j := range numbers[0] {
		var t int
		if signs[j] == "+" {
			t = 0
		} else {
			t = 1
		}
		for i := range numbers {
			n, _ := strconv.Atoi(numbers[i][j])
			if signs[j] == "+" {
				t += n
			} else {
				t *= n
			}
		}
		answers = append(answers, t)
	}
	sum := 0
	for _, v := range answers {
		sum += v
	}
	fmt.Println(sum)
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	numbers := lines[:len(lines)-1]
	signs := lines[len(lines)-1:][0]

	var indexs []int
	for i, r := range signs {
		s := string(r)
		if s != " " {
			indexs = append(indexs, i)
		}
	}

	var moreNums [][]string
	for i := 0; i < len(indexs); i++ {
		var l []string
		start := indexs[i]
		var end int
		if i+1 == len(indexs) {
			end = len(numbers[0])
		} else {
			end = indexs[i+1]
		}

		for _, n := range numbers {
			l = append(l, n[start:end])
		}

		moreNums = append(moreNums, l)
	}

	signFields := strings.Fields(signs)
	var answers []int
	for signIndex, n := range moreNums {
		var t int
		if signFields[signIndex] == "+" {
			t = 0
		} else {
			t = 1
		}

		ml := 0

		for _, s := range n {
			if len(s) > ml {
				ml = len(s)
			}
		}

		var nnn []string
		for i := ml - 1; i >= 0; i-- {
			ss := ""
			for _, nn := range n {
				if string(nn[i]) != " " {
					ss += string(nn[i])
				}
			}
			if ss != "" {

				nnn = append(nnn, ss)
			}
		}
		for _, v := range nnn {
			ii, _ := strconv.Atoi(v)
			if signFields[signIndex] == "+" {
				t += ii
			} else {
				t *= ii
			}
		}
		answers = append(answers, t)
	}
	sum := 0
	for _, v := range answers {
		sum += v
	}
	fmt.Println(sum)
}
