package problems

import "unicode"

func IsValidString(str string) bool {

	for i := 1; i < len(str); i++ {

		if str[i-1] == '+' && str[i+1] == '+' && unicode.IsLetter(rune(str[i])) {
			return true
		}
	}
	return false
}
