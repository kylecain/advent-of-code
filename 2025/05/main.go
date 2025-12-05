package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ranges [][]int
	var ids []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			var l, u int
			fmt.Sscanf(line, "%d-%d", &l, &u)
			ranges = append(ranges, []int{l, u})
		} else if len(line) > 0 {
			i, _ := strconv.Atoi(line)
			ids = append(ids, i)
		}
	}

	fmt.Println(part1(ranges, ids))
	fmt.Println(part2(ranges))
}

func part1(ranges [][]int, ids []int) int {
	sum := 0

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				sum++
				break
			}
		}
	}

	return sum
}

func part2(ranges [][]int) int {
	sum := 0

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i+1][0] <= ranges[i][1] {
			ranges[i+1][0] = min(ranges[i][0], ranges[i+1][0])
			ranges[i+1][1] = max(ranges[i][1], ranges[i+1][1])
			ranges = slices.Delete(ranges, i, i+1)
			i--
		}

	}

	for _, r := range ranges {
		sum += r[1] - r[0] + 1
	}

	return sum
}
