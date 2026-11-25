package String_problems

func ReverseString(str string) string {
	res := []rune{}
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {

		res = append(res, rune(str[i]))

	}
	return string(res)
}
