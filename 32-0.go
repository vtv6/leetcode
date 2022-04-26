package main

func longestValidParentheses(s string) int {
	result := 0
	dp := make([]int, len(s))

	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}

			} else {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			if result < dp[i] {
				result = dp[i]
			}
		}
	}
	return result
}
