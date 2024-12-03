package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

    firstList := make([]int, 0)
    secondList := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		firstList = append(firstList, first)
		secondList = append(secondList, second)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)
    total := 0
    for i, _ := range firstList {
        total += int(math.Abs(float64(firstList[i] - secondList[i])))
    }

    fmt.Printf("Total distances: %d\n", total)
}
