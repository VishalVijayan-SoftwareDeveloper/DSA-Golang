package map_problems

import (
	"fmt"
)

func CountFrequency(nums []int) map[int]int {

	count := make(map[int]int)
	for _, val := range nums {
		count[val]++
	}
	fmt.Println(count)
	return count
}

func ContainsDuplicate(nums []int) bool {

	/*	seen := make(map[int]bool)

		for _, val := range nums {
			if seen[val] {
				return true
			}
			seen[val] = true
		}
		return false*/
	seen := make(map[int]struct{})

	for _, val := range nums {
		seen[val] = struct{}{}
		fmt.Println(seen)
		if _, exist := seen[val]; exist {
			return true
		}
	}
	return false

}

// complexity: O(n) and auxilarry space: O(n)
func TwoSumPractise(input []int, target int) []int {

	seen := make(map[int]int)

	for i, val := range input {
		difference := target - val
		if index, exist := seen[difference]; exist {
			return []int{index, i}
		}
		seen[val] = i
	}
	return []int{}
}

func FirstUnique(nums []int) int {

	seen := make(map[int]int)

	for _, val := range nums {

		seen[val]++
	}
	for _, val := range nums {
		if seen[val] == 1 {
			return val
		}
	}
	return 0
}
