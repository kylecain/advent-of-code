package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	LightDiagram []bool
	Buttons      []Button
}

func NewMachine(diagram string, buttons []string) *Machine {
	var ld []bool
	for _, r := range diagram[1 : len(diagram)-1] {
		state := false
		if r == '#' {
			state = true
		}
		ld = append(ld, state)
	}

	var bs []Button
	for _, s := range buttons {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(s, -1)

		var nums []int
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			nums = append(nums, n)
		}
		bs = append(bs, Button{nums})
	}

	return &Machine{
		LightDiagram: ld,
		Buttons:      bs,
	}
}

type Button struct {
	LightIndicies []int
}

type Node struct {
	Lights []bool
	Count  int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var machines []Machine
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		machines = append(machines, *NewMachine(fields[0], fields[1:len(fields)-1]))
	}
	fmt.Println(part1(machines))
}

func part1(machines []Machine) int {
	sum := 0
	for _, machine := range machines {
		sum += bfs(machine)
	}
	return sum
}

func bfs(machine Machine) int {
	start := make([]bool, len(machine.LightDiagram))
	queue := []Node{{start, 0}}

	var found Node
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if slices.Equal(node.Lights, machine.LightDiagram) {
			found = node
			break
		}

		for _, b := range machine.Buttons {
			lightsCopy := make([]bool, len(node.Lights))
			copy(lightsCopy, node.Lights)

			for _, i := range b.LightIndicies {
				lightsCopy[i] = !lightsCopy[i]
			}
			queue = append(queue, Node{lightsCopy, node.Count + 1})
		}

	}

	return found.Count
}
