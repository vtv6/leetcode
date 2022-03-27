package main

var lettermap = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	return calc([]rune(digits))
}

func calc(s []rune) []string {
	result := []string{}

	if len(s) == 0 {
		return []string{}
	}

	temp := calc(s[1:])

	for _, c := range lettermap[s[0]] {

		if len(temp) == 0 {
			result = append(result, string(c))
		} else {
			for _, ss := range temp {
				result = append(result, string(c)+ss)
			}
		}
	}

	return result
}
