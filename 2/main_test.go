package main

import "testing"

type IncreaseTest struct {
	Reading     []int
	Expected    bool
	Description string
}

func (st IncreaseTest) Test() bool {
	return mustIncrease(st.Reading)
}

type DecreaseTest struct {
	Reading     []int
	Expected    bool
	Description string
}

func (st DecreaseTest) Test() bool {
	return mustDecrease(st.Reading)
}

type SafetyTest struct {
	Reading     []int
	Expected    bool
	Description string
}

func (st SafetyTest) Test() bool {
	return isSafe(st.Reading)
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
		{[]int{}, true, "Safe"},
		{[]int{}, true, "Unsafe"},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("%s - input: %+v, wanted %t", test.Description, test.Reading, test.Expected)
		}
	}
}
