package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	IsAnimal bool
	Symbol   string
	X        int
	Y        int
	Previous *Node
	Next     *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) append(isAnimal bool, symbol string, x int, y int) {
	newNode := &Node{
		IsAnimal: isAnimal,
		Symbol:   symbol,
		X:        x,
		Y:        y,
	}

	if list.Tail == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	current := list.Tail
	current.Next = newNode
	newNode.Previous = current
	list.Tail = newNode
}

func (list *LinkedList) IsConnectingSymbol(symbol string, direction string) bool {
	if symbol == "." {
		return false
	}

	var possibleSymbols []string
	if direction == "up" {
		possibleSymbols = upMap[list.Tail.Symbol]
	} else if direction == "down" {
		possibleSymbols = downMap[list.Tail.Symbol]
	} else if direction == "right" {
		possibleSymbols = rightMap[list.Tail.Symbol]
	} else if direction == "left" {
		possibleSymbols = leftMap[list.Tail.Symbol]
	}

	for _, possibleSymbol := range possibleSymbols {
		if symbol == possibleSymbol {
			return true
		}
	}

	return false
}

func (list *LinkedList) Length() int {
	current := list.Head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

var leftMap map[string][]string
var rightMap map[string][]string
var downMap map[string][]string
var upMap map[string][]string

func main() {
	leftMap = map[string][]string{
		"S": {"|", "-", "L", "J", "7", "F"},
		"|": {},
		"-": {"-", "L", "J", "7", "F", "S"},
		"L": {},
		"J": {"-", "L", "F", "S"},
		"7": {"-", "L", "F", "S"},
		"F": {},
	}
	rightMap = map[string][]string{
		"S": {"|", "-", "L", "J", "7", "F"},
		"|": {},
		"-": {"-", "L", "J", "7", "F", "S"},
		"L": {"-", "J", "7", "S"},
		"J": {},
		"7": {},
		"F": {"-", "J", "7", "S"},
	}

	upMap = map[string][]string{
		"S": {"|", "-", "L", "J", "7", "F"},
		"|": {"|", "7", "F", "S"},
		"-": {},
		"L": {"|", "7", "F", "S"},
		"J": {"|", "7", "F", "S"},
		"7": {},
		"F": {},
	}

	downMap = map[string][]string{
		"S": {"|", "-", "L", "J", "7", "F"},
		"|": {"|", "L", "J", "7", "F", "S"},
		"-": {},
		"L": {},
		"J": {},
		"7": {"|", "L", "J", "S"},
		"F": {"|", "L", "J", "S"},
	}

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	animalFound := false
	linkedList := &LinkedList{}

	for i, str := range data {
		if animalFound {
			break
		}

		for j, sym := range str {
			if s := string(sym); s != "." {
				if s == "S" {
					linkedList.append(true, "S", j, i)
					animalFound = true
					break
				}
			}
		}
	}

	i := 0

	for {
		if i != 0 && linkedList.Tail.Symbol == "S" {
			break
		}

		x, y := linkedList.Tail.X, linkedList.Tail.Y
		previousX, previousY := 0, 0

		if linkedList.Tail.Previous != nil {
			previousX = linkedList.Tail.Previous.X
			previousY = linkedList.Tail.Previous.Y
		}

		if x-1 >= 0 {
			symbol := string(data[y][x-1])

			if x-1 != previousX || y != previousY {
				if linkedList.IsConnectingSymbol(symbol, "left") {
					linkedList.append(false, symbol, x-1, y)
					i++
					continue
				}
			}
		}

		if x+1 < len(data[0]) {
			symbol := string(data[y][x+1])

			if x+1 != previousX || y != previousY {
				if linkedList.IsConnectingSymbol(symbol, "right") {
					linkedList.append(false, symbol, x+1, y)
					i++
					continue
				}
			}
		}

		if y-1 >= 0 {
			symbol := string(data[y-1][x])

			if x != previousX || y-1 != previousY {
				if linkedList.IsConnectingSymbol(symbol, "up") {
					linkedList.append(false, symbol, x, y-1)
					i++
					continue
				}
			}
		}

		if y+1 < len(data) {
			symbol := string(data[y+1][x])

			if x != previousX || y+1 != previousY {
				if linkedList.IsConnectingSymbol(symbol, "down") {
					linkedList.append(false, symbol, x, y+1)
					i++
					continue
				}
			}
		}
	}

	fmt.Println(linkedList.Length()/2 - 1)

	current := linkedList.Head
	sum := 0
	for current.Next != nil {
		left := current.X * current.Next.Y
		right := current.Y * current.Next.X
		fmt.Println(left, " - ", right)
		sum += left - right
		current = current.Next
	}
	fmt.Println(int(math.Abs(float64(sum))/2) - (linkedList.Length()/2 - 1))
}
