package main

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = len(dp)
	}
	dp[0] = 0

	for i := 0; i < len(dp)-1; i++ {
		for step := 1; step <= min(nums[i], len(dp)-i-1); step++ {
			dp[i+step] = min(dp[i+step], dp[i]+1)
		}
	}
	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
