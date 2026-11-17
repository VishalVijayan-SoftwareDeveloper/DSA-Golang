/*nums := []int{2, 7, 11, 15}
target := 9
output: []int{0, 1}*/

package map_problems

func TwoSum(input []int, target int) []int {

	seen := make(map[int]int)

	for i := 0; i < len(input); i++ {
		compliment := target - input[i]
		if j, ok := seen[compliment]; ok {
			return []int{j, i}
		}
		seen[input[i]] = i

	}
	return []int{}
}
