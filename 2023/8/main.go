package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Root  string
	Left  string
	Right string
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	var nodes []Node

	instructions := ""
	i := 0

	for scanner.Scan() {
		if text := scanner.Text(); text == "" {
			continue
		} else if i == 0 {
			instructions = text
		} else {
			strippedText := reg.ReplaceAllString(text, " ")
			splitString := strings.Fields(strippedText)

			node := Node{
				Root:  splitString[0],
				Left:  splitString[1],
				Right: splitString[2],
			}

			nodes = append(nodes, node)
		}

		i++
	}

	nodesMap := make(map[string][]string)
	var nodesEndingInA []string

	for _, node := range nodes {
		nodesMap[node.Root] = []string{node.Left, node.Right}

		lastChar := string(node.Root[len(node.Root)-1])
		if lastChar == "A" {
			nodesEndingInA = append(nodesEndingInA, node.Root)
		}
	}

	var instructionsInCycle []int64

	for _, node := range nodesEndingInA {
		currentNode := node
		shouldSearch := true
		j := 0
		for shouldSearch {
			for _, instruction := range instructions {
				lastChar := string(currentNode[len(currentNode)-1])

				if lastChar == "Z" {
					shouldSearch = false
					instructionsInCycle = append(instructionsInCycle, int64(j))
					break
				}

				str := string(instruction)
				if str == "L" {
					currentNode = nodesMap[currentNode][0]
				} else {
					currentNode = nodesMap[currentNode][1]
				}

				j++
			}
		}
	}

	fmt.Println(findLCM(instructionsInCycle))
}

// Function to calculate the greatest common divisor (GCD) of two numbers
func gcd(a, b int64) int64 {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	var result big.Int
	return result.GCD(nil, nil, bigA, bigB).Int64()
}

// Function to calculate the least common multiple (LCM) of two numbers
func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}

// Function to find the LCM of integers in an array
func findLCM(arr []int64) int64 {
	if len(arr) == 0 {
		return 0
	}
	lcmValue := arr[0]
	for i := 1; i < len(arr); i++ {
		lcmValue = lcm(lcmValue, arr[i])
	}
	return lcmValue
}
