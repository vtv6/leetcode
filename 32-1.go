package main

func longestValidParentheses(s string) int {
	stack := []int{-1}
	result := 0

	for i := range s {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			if s[i] == '(' {
				stack = append(stack, i)
			} else {
				lastIndex := len(stack) - 1
				if stack[lastIndex] != -1 && s[stack[lastIndex]] == '(' {
					if i-stack[lastIndex-1] > result {
						result = i - stack[lastIndex-1]
					}
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, i)
				}
			}
		}
	}
	return result
}
