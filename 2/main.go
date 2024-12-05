package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

// mustIncrease returns true when all values in the
// slice are increasing. All values must increase by
// at least 1, but no more than 3.
func mustIncrease(list []int) bool {
	last := -1
	for _, v := range list {
		if last == -1 || last < v && last+3 >= v {
			last = v
			continue
		} else {
			return false
		}
	}
	return true
}

// mustDecrease returns true when all values in the
// slice are decreasing. All values must decrease by
// at least 1, but no more than 3.
func mustDecrease(list []int) bool {
	last := -1
	for _, v := range list {
		if last == -1 || last > v && last-3 <= v {
			last = v
			continue
		} else {
			return false
		}
	}
	return true
}

// isSafe checks to see if a reading (list) is either
// all increasing or all decreasing.
func isSafe(list []int) bool {
	return mustDecrease(list) || mustIncrease(list)
}

// parseReading will take the raw string from the readings
// output and returns a properly formatted reading.
func parseReading(input string) []int {
	reading := make([]int, 0)
	raw := strings.Split(input, " ")
	for _, v := range raw {
		vv, _ := strconv.Atoi(v)
		reading = append(reading, vv)
	}
	return reading
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	safeReportCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reading := parseReading(scanner.Text())
		if isSafe(reading) {
			safeReportCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	fmt.Printf("Safe reports: %d\n", safeReportCount)
}
