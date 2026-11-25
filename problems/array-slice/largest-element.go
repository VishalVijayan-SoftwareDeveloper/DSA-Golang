package array_slice

func LargestElement(input []int) int {

	max := input[0]
	for _, value := range input {
		if value > max {
			max = value
		}
	}
	return max
}

func SumArray(nums []int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}
	return sum
}

func Contains(nums []int, target int) bool {

	for _, val := range nums {
		if target == val {
			return true
		}
	}
	return false
}
