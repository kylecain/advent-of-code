package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var devices [][]string
	for scanner.Scan() {
		re := regexp.MustCompile(`\w+`)
		matches := re.FindAllString(scanner.Text(), -1)
		devices = append(devices, matches)

	}

	deviceMap := make(map[string][]string)
	for _, d := range devices {
		deviceMap[d[0]] = d[1:]
	}

	// fmt.Println(part1(deviceMap))
	fmt.Println(part2(deviceMap))
}

func part1(deviceMap map[string][]string) int {
	count := 0
	dfs(deviceMap, "you", &count)
	return count
}

func part2(deviceMap map[string][]string) int {
	count := 0
	memo := make(map[string]bool)
	dfs2(deviceMap, memo, "svr", &count, false, false)
	return count
}

func dfs(deviceMap map[string][]string, current string, count *int) {
	if current == "out" {
		*count++
		return
	}

	for _, v := range deviceMap[current] {
		dfs(deviceMap, v, count)
	}
}

func dfs2(deviceMap map[string][]string, memo map[string]bool, current string, count *int, seenFft, seenDac bool) {
	if v, ok := memo[current]; ok {
		if v {
			*count++
		}
		return
	}

	if current == "out" && seenFft && seenDac {
		*count++
		memo[current] = true
		return
	} else if current == "out" {
		memo[current] = false
		return
	}

	for _, v := range deviceMap[current] {
		if v == "fft" {
			seenFft = true
		}
		if v == "dac" {
			seenDac = true
		}
		dfs2(deviceMap, memo, v, count, seenFft, seenDac)
	}
}
