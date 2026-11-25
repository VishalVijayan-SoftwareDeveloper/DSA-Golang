package array_slice

import (
	"fmt"
	"testing"
)

func TestLargestElement(t *testing.T) {

	/*Max := LargestElement([]int{1, 2, 3, 4})
	if Max != 4 {
		t.Errorf("LargestElement(#{[]int{1,2,3,4}}) = #{Max}")
	}*/

	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{-1, -2, -3, -4}, -1},
		{[]int{1, 7, 3, 6}, 7},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.input)
		t.Run(testName, func(t *testing.T) {
			max := LargestElement(tt.input)
			if max != tt.want {
				t.Errorf("LargestElement(%v) = %d, want %d",
					tt.input, max, tt.want)
			}
		})
	}
}

func TestSumArray(t *testing.T) {

	/*Max := LargestElement([]int{1, 2, 3, 4})
	if Max != 4 {
		t.Errorf("LargestElement(#{[]int{1,2,3,4}}) = #{Max}")
	}*/

	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-1, -2, -3, -4}, -10},
		{[]int{1, 7, 3, 6}, 17},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.input)
		t.Run(testName, func(t *testing.T) {
			max := SumArray(tt.input)
			if max != tt.want {
				t.Errorf("SumArray(%v) = %d, want %d",
					tt.input, max, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {

	tests := []struct {
		input  []int
		target int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 5, true},
		{[]int{11, 12, 3, 4, 15, 26}, 26, true},
		{[]int{1, 2, 3, 4, 5, 6}, 7, false},
	}

	for _, tt := range tests {

		t.Run("Tests", func(t *testing.T) {
			got := Contains(tt.input, tt.target)
			if got != tt.want {
				t.Errorf("Result is wrong")
			}
		})
	}
}
