package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}
}
