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

// countOccurrences will return the number of times that v
// is found in list where list is sorted.
func countOccurrences[V comparable](value V, list []V) int {
	total := 0
	found := false
	for _, i := range list {
		if value == i {
			found = true
			total += 1
		} else {
			// We can break early because the list is sorted.
			if found {
				break
			}
		}
	}
	return total
}

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

	matches := make(map[int]int, 0)
	for _, v := range firstList {
		// Find all numbers in second list and calculate
		// similarity score.
		if v, ok := matches[v]; !ok {
			matches[v] = 0
		}

		// Scan secondList for occurrences of v...
		foundMultiplier := countOccurrences(v, secondList)

		// Add multiplier for later similarity calculation...
		matches[v] += foundMultiplier

	}

	fmt.Printf("Total distances: %d\n", total)
	similarity := 0
	for num, mult := range matches {
		similarity += num * mult
	}
	fmt.Printf("Similarity score: %d\n", similarity)
}
