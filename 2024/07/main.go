package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type CalibrationEquation struct {
    total int
    values []int
}

type Node struct {
    value int
    sign string
    visited bool
    children []*Node
}


var signs = []string{"+", "*", "||"}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    calibrationEquations := []CalibrationEquation{}

	for scanner.Scan() {
		line := scanner.Text()

		exp := regexp.MustCompile(`\d+`)
		numbers := exp.FindAllString(line, -1)

        total, _ := strconv.Atoi(numbers[0])

        values := []int{}

        for _, num := range numbers[1:] {
            n, _ := strconv.Atoi(num)
            values = append(values, n)
        }

        calibrationEquations = append(calibrationEquations, CalibrationEquation{
            total,
            values,
        })
	}


	// partOne(calibrationEquations)
    partTwo(calibrationEquations)
}

func partOne(calibrationEquations []CalibrationEquation) {
    sum := 0
    for _, calcalibrationEquation := range calibrationEquations {
        root := buildGraph(calcalibrationEquation)
        isSolved := isSolveable(calcalibrationEquation.total, root.value, root)
        if isSolved {
            sum += calcalibrationEquation.total
        }
    }
    fmt.Println(sum)
}

func partTwo(calibrationEquations []CalibrationEquation) {
    sum := 0
    for _, calcalibrationEquation := range calibrationEquations {
        root := buildGraph(calcalibrationEquation)
        isSolved := isSolveable(calcalibrationEquation.total, root.value, root)
        if isSolved {
            sum += calcalibrationEquation.total
        }
    }
    fmt.Println(sum)
}

func buildGraph(calibrationEquation CalibrationEquation) *Node {
    root := &Node{calibrationEquation.values[0], "+", false, []*Node{}}

    for i, value := range calibrationEquation.values {
        if i == 0 { continue }

        for _, node := range root.findChildlessNodes() {
            node.addChildrenWithValue(value)
        }
    }

    return root
}

func (node *Node) findChildlessNodes() []*Node {
    childlessNodes := []*Node{}
    stack := []*Node{}

    stack = append(stack, node)

    for len(stack) > 0 {
        currentNode := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]

        if len(currentNode.children) == 0 {
            childlessNodes = append(childlessNodes, currentNode)
        }

        for _, n := range currentNode.children {
            stack = append(stack, n)
        }
    }

    return childlessNodes
}

func (node *Node) addChildrenWithValue(newValue int) {
    for _, sign := range signs {
        node.children = append(node.children, &Node{
            newValue,
            sign,
            false,
            []*Node{}},
        )
    }
}

func isSolveable(target int, total int, node *Node) bool {
    if total == target && len(node.children) == 0 {
        return true
    }

    if node == nil || node.visited {
        return false
    }

    node.visited = true

    for _, n := range node.children {
        if n.sign == "+" {
            if isSolveable(target, total + n.value, n) {
                return true
            }
        }

        if n.sign == "*" {
            if isSolveable(target, total * n.value, n) {
                return true
            }
        }

        if n.sign == "||" {
            a, b := strconv.Itoa(total), strconv.Itoa(n.value)
            c, _ := strconv.Atoi(a+b)
            if isSolveable(target, c, n) {
                return true
            }
        }
    }

    return false
}

