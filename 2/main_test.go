// Template file only!
// Saving some time each day to ensure I
// actually write tests.
package main

import "testing"

type IncreaseTest struct {
	Reading  []int
	Expected bool
}

func (st IncreaseTest) Test() bool {
	return mustIncrease(st.Reading)
}

type DecreaseTest struct {
	Reading  []int
	Expected bool
}

func (st DecreaseTest) Test() bool {
	return mustDecrease(st.Reading)
}

func TestMustIncrease(t *testing.T) {
	tests := []IncreaseTest{
		{[]int{}, true},
		{[]int{}, false},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("input: %+v - wanted %t", test.Reading, test.Expected)
		}
	}
}

func TestMustDecrease(t *testing.T) {
	tests := []DecreaseTest{
		{[]int{}, true},
		{[]int{}, false},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("input: %+v - wanted %t", test.Reading, test.Expected)
		}
	}
}
