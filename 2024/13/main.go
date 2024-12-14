package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Machine struct {
    eq1 []int
    eq2 []int
}

func main() {
    data, _ := os.ReadFile("input.txt")
    exp := regexp.MustCompile(`\d+`)
    buttons := exp.FindAllString(string(data), -1)

    machines := []Machine{}
    var machine Machine

    for i, button := range buttons {
        v, _ := strconv.Atoi(button)
        if len(machine.eq2) == 3 {
            machines = append(machines, machine)
            machine = Machine{}
        }

        if i % 2 == 0 {
            if len(machine.eq1) == 2 {
                machine.eq1 = append(machine.eq1, v + 10000000000000)
            } else {
                machine.eq1 = append(machine.eq1, v)
            }
        } else {
            if len(machine.eq2) == 2 {
                machine.eq2 = append(machine.eq2, v + 10000000000000)
            } else {
                machine.eq2 = append(machine.eq2, v)
            }
        }
    }
    machines = append(machines, machine)

    partOne(machines)
}

func partOne(machines []Machine) {
    outerTotal := 0

    for _, machine := range machines {
        orig := make([]int, len(machine.eq1))
        copy(orig, machine.eq1)

        t1, t2 := machine.eq1[0], machine.eq2[0]
        for i := 0; i < len(machine.eq1); i++ {
            machine.eq1[i] = machine.eq1[i] * t2
        }

        for i := 0; i < len(machine.eq2); i++ {
            machine.eq2[i] = machine.eq2[i] * t1
        }

        for i := 0; i < len(machine.eq2); i++ {
            machine.eq2[i] = machine.eq2[i] - machine.eq1[i]
        }

        if !(machine.eq2[2] % machine.eq2[1] == 0) {
            continue
        }

        b := machine.eq2[2] / machine.eq2[1]

        fmt.Println(b)
        fmt.Println(orig)

        if !(((orig[2] - (orig[1] * b)) % orig[0]) == 0) {
            continue
        }

        fmt.Println("im here")
        a := 3 * ((orig[2] - (orig[1] * b)) / orig[0])

        fmt.Println(a)
        innerTotal := a + b
        outerTotal += innerTotal
    }

    fmt.Println("outer total", outerTotal)
}

