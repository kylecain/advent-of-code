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
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()

        for _, v := range strings.Fields(line) {
            n, _ := strconv.Atoi(v)
            numbers = append(numbers, n)       
        }
	}


	// partOne(numbers)
	partTwo(numbers)
}

func partOne(numbers []int) {
    fmt.Println(numbers)
    // order matters btw

    for j := 0; j < 25; j++ {
        for i := 0; i < len(numbers); i++ {
            if numbers[i] == 0 {
                numbers[i]++
            } else if str := strconv.Itoa(numbers[i]); len(str) % 2 == 0 {
                mid := len(str) / 2
                leftStr, rightStr := str[:mid], str[mid:]
                leftInt, _ := strconv.Atoi(leftStr)
                rightInt, _ := strconv.Atoi(rightStr)

                temp := []int{}
                temp = append(temp, numbers[:i]...)
                temp = append(temp, leftInt)
                temp = append(temp, rightInt)
                temp = append(temp,  numbers[i+1:]...)

                numbers = temp
                i++
            } else {
                numbers[i] = numbers[i] * 2024
            }
        }
    }

    fmt.Println(len(numbers))
}


func partTwo(numbers []int) {
    m := make(map[int]int)
    
    for _, v := range numbers {
        m[v]++
    }

    for j := 0; j < 75; j++ {
        m1 := make(map[int]int)
        
        for k, v := range m {
            if k == 0 {
                m1[1] += v
            } else if str := strconv.Itoa(k); len(str) % 2 == 0 {
                mid := len(str) / 2
                leftStr, rightStr := str[:mid], str[mid:]
                leftInt, _ := strconv.Atoi(leftStr)
                rightInt, _ := strconv.Atoi(rightStr)

                m1[leftInt] += v
                m1[rightInt] += v
            } else {
                m1[k * 2024] += v
            }
        }

        m = m1
    }

    sum := 0 
    for _, v := range m {
        sum += v
    }

    fmt.Println(sum)
}
