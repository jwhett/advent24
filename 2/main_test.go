package main

import (
	"slices"
	"testing"
)

type IncreaseTest struct {
	Reading     []int
	Expected    bool
	Description string
}

func (it IncreaseTest) Test() bool {
	return mustIncrease(it.Reading)
}

type DecreaseTest struct {
	Reading     []int
	Expected    bool
	Description string
}

func (dt DecreaseTest) Test() bool {
	return mustDecrease(dt.Reading)
}

type SafetyTest struct {
	Reading     []int
	Expected    bool
	Description string
}

func (st SafetyTest) Test() bool {
	return isSafe(st.Reading)
}

type ParserTest struct {
	Input       string
	Expected    []int
	Description string
}

func (pt ParserTest) Test() bool {
	return slices.Compare(parseReading(pt.Input), pt.Expected) == 0
}

func TestMustIncrease(t *testing.T) {
	tests := []IncreaseTest{
		{[]int{1, 3, 6, 7, 9}, true, "Stable increase"},
		{[]int{1, 3, 6, 6, 7}, false, "Unstable increase; doubles"},
		{[]int{1, 3, 6, 5, 7}, false, "Unstable increase; increase and decrease"},
		{[]int{7, 6, 4, 2, 1}, false, "Decrease"},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("input: %+v - wanted %t", test.Reading, test.Expected)
		}
	}
}

func TestMustDecrease(t *testing.T) {
	tests := []DecreaseTest{
		{[]int{7, 6, 4, 2, 1}, true, "Stable decrease"},
		{[]int{7, 6, 6, 4, 2}, false, "Unstable decrease; doubles"},
		{[]int{7, 6, 7, 4, 2}, false, "Unstable decrease; increase and decrease"},
		{[]int{1, 2, 7, 8, 9}, false, "Increase"},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("%s - input: %+v, wanted %t", test.Description, test.Reading, test.Expected)
		}
	}
}

func TestSafety(t *testing.T) {
	tests := []SafetyTest{
		{[]int{7, 6, 4, 2, 1}, true, "Safe, decreasing"},
		{[]int{1, 2, 7, 8, 9}, false, "Unsafe, too great of an increase"},
		{[]int{9, 7, 6, 2, 1}, false, "Unsafe, too great of a decrease"},
		{[]int{1, 3, 2, 4, 5}, false, "Unsafe, increasing and decreasing"},
		{[]int{8, 6, 4, 4, 1}, false, "Unsafe, double readings"},
		{[]int{1, 3, 6, 7, 9}, true, "Safe, increasing"},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("%s - input: %+v, wanted %t", test.Description, test.Reading, test.Expected)
		}
	}
}

func TestParser(t *testing.T) {
	tests := []ParserTest{
		{"7 6 4 2 1", []int{7, 6, 4, 2, 1}, "Safe, decreasing"},
	}

	for _, test := range tests {
		if !test.Test() {
			t.Errorf("%s - input: %+v, wanted %v", test.Description, test.Input, test.Expected)
		}
	}
}
