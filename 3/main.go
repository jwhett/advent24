package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

func findInstructions(input string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return re.FindAllString(input, -1)
}

func parseMul(input string) int {
	re := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	parsed := re.FindString(input)
	rawNumbers := strings.Split(parsed, ",")
	numbers := make([]int, 0)
	for _, v := range rawNumbers {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)
	}
	return numbers[0] * numbers[1]
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	found := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		found = append(found, findInstructions(line)...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	sum := 0
	for _, mul := range found {
		sum += parseMul(mul)
	}

	fmt.Printf("Sum of all products: %d\n", sum)
}
