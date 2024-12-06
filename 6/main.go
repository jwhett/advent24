package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// Guard Orientations
const (
	UP    = '^'
	DOWN  = 'v'
	LEFT  = '<'
	RIGHT = '>'
)

// Obsticals
const (
	BLOCKER = '#'
)

const (
	inputFile = "input"
)

type Position struct {
	X, Y int
}

type Board struct {
	Rows [][]rune
	Position
}

// scanFor searches the board for an arbitrary target
// rune and return its position. Returns Position{-1, -1}
// and an error when the target is not found.
func (b *Board) scanFor(target rune) (Position, error) {
	for i, row := range b.Rows {
		if found := slices.Index(row, target); found > -1 {
			return Position{i, found}, nil
		}
	}
	return Position{-1, -1}, fmt.Errorf("couldn't find target: %v", target)
}

// scanForGuard returns the postition of the guard
// on the board. Returns an invalid position and an error
// in the case that the guard is no longer on the map.
func (b *Board) scanForGuard() (Position, error) {
	// TODO: Scan the board and return the first
	// Position of target.
	guards := []rune{UP, DOWN, LEFT, RIGHT}
	for _, guard := range guards {
		if position, err := b.scanFor(guard); err == nil {
			return position, nil
		}
	}
	return Position{-1, -1}, fmt.Errorf("couldn't find the guard")
}

// findAndSetGuardPosition updates the Board with
// the position of the guard after searching the board.
func (b *Board) findAndSetGuardPosition() error {
	var err error
	if b.Position, err = b.scanForGuard(); err != nil {
		return err
	}
	return nil
}

// NewBoard returns a fully built Board along with
// the guard's current position.
func NewBoard(in []string) Board {
	board := Board{}
	for _, line := range in {
		runes := make([]rune, 0)
		for _, r := range line {
			runes = append(runes, r)
		}
		board.Rows = append(board.Rows, runes)
	}
	board.findAndSetGuardPosition()
	return board
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	board := NewBoard(lines)
	fmt.Printf("Rows: %d, Row length: %d, Guard's position: %v\n", len(board.Rows), len(board.Rows[0]), board.Position)
}
