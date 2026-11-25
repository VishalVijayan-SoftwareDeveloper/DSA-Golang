package map_problems

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountFrequency(t *testing.T) {
	nums := []int{1, 2, 2, 3, 1, 4, 2, 3}
	res := CountFrequency(nums)
	want := map[int]int{
		1: 2,
		2: 3,
		3: 2,
		4: 1,
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("Map are not equal, failed")
	}
}
func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 2, 3, 1}
	res := ContainsDuplicate(nums)
	if !res {
		t.Errorf("Test Failed")
	}
}
func TestTwoSumPractise(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output []int
	}{
		{[]int{1, 2, 3, 5}, 3, []int{0, 1}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {

		testName := fmt.Sprintf("test")

		t.Run(testName, func(t *testing.T) {
			get := TwoSumPractise(tt.input, tt.target)

			if !reflect.DeepEqual(get, tt.output) {
				t.Errorf("test failed, wrong output")
			}
		})

	}
}

func TestFirstUnique(t *testing.T) {
	res := FirstUnique([]int{4, 5, 1, 2, 1, 4, 8, 5})
	if res != 2 {
		t.Errorf("test failed, wrong output")
	}
}
