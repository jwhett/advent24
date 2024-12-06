package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

type ErrorTracker struct {
	IncreaseErrors, DecreaseErrors int
}

// mustIncrease returns true when all values in the
// slice are increasing. All values must increase by
// at least 1, but no more than 3.
func mustIncrease(tracker ErrorTracker, list []int) bool {
	last := -1
	for i, v := range list {
		if last == -1 || last < v && last+3 >= v {
			last = v
			continue
		} else {
			tracker.IncreaseErrors += 1
			if tracker.IncreaseErrors > 1 {
				return false
			}
			newList := slices.Delete(list, i, i+1)
			mustIncrease(tracker, newList)
		}
	}
	return true
}

// mustDecrease returns true when all values in the
// slice are decreasing. All values must decrease by
// at least 1, but no more than 3.
func mustDecrease(tracker ErrorTracker, list []int) bool {
	last := -1
	for i, v := range list {
		if last == -1 || last > v && last-3 <= v {
			last = v
			continue
		} else {
			tracker.DecreaseErrors += 1
			if tracker.DecreaseErrors > 1 {
				return false
			}
			newList := slices.Delete(list, i, i+1)
			mustDecrease(tracker, newList)
		}
	}
	return true
}

// isSafe checks to see if a reading (list) is either
// all increasing or all decreasing.
func isSafe(list []int) bool {
	return mustDecrease(ErrorTracker{}, list) || mustIncrease(ErrorTracker{}, list)
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
