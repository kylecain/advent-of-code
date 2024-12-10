package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
    data := []string{}

	for scanner.Scan() {
		line := scanner.Text()
        data = parser(line)

	}

	// partOne(data)
	partTwo(data)
}

func parser(line string) []string {
    data := []string{}
    isDot := false
    id := 0

    for _, v := range line {
        numAsInt, _ := strconv.Atoi(string(v))
        idAsString := strconv.Itoa(id)

        if isDot {
            for i := 0; i < numAsInt; i++ {
                data = append(data, ".")
            }
        } else {
            for i := 0; i < numAsInt; i++ {
                data = append(data, idAsString)
            }

            id++
        }

        isDot = !isDot
    }

    return data
}

func partOne(data []string) {
    isSorted := false
    for i := 0; i < len(data); i++ {
        if isSorted { break }

        if data[i] == "." {
            for j := len(data) - 1; j >= 0; j-- {
                if j == i {
                    isSorted = true
                    break
                }
                
                if data[j] != "." {
                    data[i], data[j] = data[j], data[i]
                    break
                }
            }

        }
    }

    sum := 0
    for i, v := range data {
        stringAsInt, _ := strconv.Atoi(v)
        sum += i * stringAsInt
    }

    // fmt.Println(data)
    // fmt.Println(sum)
}


func partTwo(data []string) {
    for i := len(data) - 1; i >= 0; {
        neededSpace := 0
        if data[i] != "." {
            neededSpace++

            for true {
                if i - neededSpace < 0 {
                    break
                }

                if data[i - neededSpace] == data[i] {
                    neededSpace++
                } else {
                    break
                }
            }
        }

        replacementFound := false
        foundSpace := 0
        var foundSpaceStartingIndex int


        for j := 0; j < len(data[:i]); j++ {
            if data[j] == "." {
                foundSpace++
            } else if foundSpace >= neededSpace {
                replacementFound = true
                foundSpaceStartingIndex = j - foundSpace
                break
            } else {
                foundSpace = 0
            }
        }

        if replacementFound {
            for k := 0; k < neededSpace; k++ {
                data[foundSpaceStartingIndex + k], data[i - k] = data[i - k], data[foundSpaceStartingIndex + k]
            }

        } 

        if neededSpace == 0 {
            i--
        } else {
            i = i - neededSpace
        }
    }

    sum := 0
    for i, v := range data {
        stringAsInt, _ := strconv.Atoi(v)
        sum += i * stringAsInt
    }

    fmt.Println(sum)
}

