/*
Input:  "swiss"
Output: "w"

Input:  "leetcode"
Output: "l"

Input:  "aabb"
Output: ""
*/

package String_problems

func FirstNonRepeatingLetter(str string) string {

	repeatCount := make(map[rune]int)
	for _, ch := range str {
		repeatCount[ch]++
	}

	for _, ch := range str {
		if repeatCount[ch] == 1 {
			return string(ch)
		}
	}

	return ""
}

func FirstNonRepeatingLetterIndex(str string) int {

	repeatCount := make(map[rune]int)
	for _, ch := range str {
		repeatCount[ch]++
	}

	for i, ch := range str {
		if repeatCount[ch] == 1 {
			return i
		}
	}

	return -1
}

func FirstRepeatingCharacter(str string) string {
	repeatCount := make(map[rune]bool)
	for _, ch := range str {
		if repeatCount[ch] {
			return string(ch)
		} else {
			repeatCount[ch] = true
		}
	}
	return ""

}
