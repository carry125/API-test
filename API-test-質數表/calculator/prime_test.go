package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockCalculate(numStart interface{}, numEnd interface{}) []int {
	start, startOk := numStart.(int)
	end, endOk := numEnd.(int)

	if !startOk || !endOk {
		return []int{}
	}

	var numberGroup []int
	if start > end {
		start, end = Swap(start, end)
	}
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			numberGroup = append(numberGroup, i)
		}
	}
	return numberGroup
}
func TestCalculate(t *testing.T) {
	testCases := []struct {
		name     string
		input1   interface{}
		input2   interface{}
		expected []int
	}{
		{"Integer input", 1, 10, []int{2, 3, 5, 7}},
		{"Integer input", 13, 10, []int{11, 13}},
		{"Integer input", 03, 30, []int{3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{"Integer input", 1, 100000, Calculate(1, 100000)},
		{"Integer input", 2, 2, []int{2}},
		{"Integer input", 0, 0, []int{}},
		{"Non-integer input", 10, 10.5, []int{}},
		{"Non-integer input", -10, 10.5, []int{}},
		{"invalid input", "#", 10, []int{}},
		{"invalid input", " ", 10, []int{}},
		{"invalid input", "abc", "def", []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MockCalculate(tc.input1, tc.input2)
			assert.ElementsMatch(t, tc.expected, result)
		})
	}
}
