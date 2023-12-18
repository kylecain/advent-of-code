package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		numStrs := strings.Fields(scanner.Text())
		var nums []int

		for _, str := range numStrs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}

		nums2d := [][]int{nums}
		i := 0
		shouldAdd := true

		for shouldAdd {
			var subArr []int
			for j := 0; j < len(nums2d[i])-1; j++ {
				subArr = append(subArr, nums2d[i][j+1]-nums2d[i][j])
			}
			nums2d = append(nums2d, subArr)
			shouldAdd = !allZeros(subArr)
			i++
		}

		for j := len(nums2d) - 1; j > 0; j-- {
			if j == len(nums2d)-1 {
				nums2d[j] = append(nums2d[j], 0)
			}

			expectedNum := nums2d[j][len(nums2d[j])-1] + nums2d[j-1][len(nums2d[j-1])-1]
			nums2d[j-1] = append(nums2d[j-1], expectedNum)
		}

		totalSum += nums2d[0][len(nums2d[0])-1]
	}

	fmt.Println(totalSum)
}

func allZeros(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}
