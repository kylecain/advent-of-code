package main

import (
	"bufio"
	"fmt"
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
	reg := regexp.MustCompile("[^a-zA-Z]+")

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

	for _, node := range nodes {
		nodesMap[node.Root] = []string{node.Left, node.Right}
	}

	shouldSearch := true
	currentNode := "AAA"

	j := 0
	for shouldSearch {
		for _, instruction := range instructions {
			if currentNode == "ZZZ" {
				shouldSearch = false
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

	fmt.Println(j)
}
