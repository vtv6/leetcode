package main

func longestValidParentheses(s string) int {
	left, right, result := 0, 0, 0

	//     left to right
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if result < 2*left {
				result = 2 * left
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	//     right to left
	for j := range s {
		i := len(s) - j - 1
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if result < 2*left {
				result = 2 * left
			}
		} else if right < left {
			left, right = 0, 0
		}
	}

	return result
}
