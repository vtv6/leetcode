package main

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	result := ""
	count := 1
	index := 1
	currentDigit := s[0]
	for ; index < len(s); index++ {
		if s[index] == currentDigit {
			count++
		} else {
			result += string(count+'0') + string(currentDigit)
			count = 1
			currentDigit = s[index]
		}
	}
	result += string(count+'0') + string(currentDigit)
	return result
}
