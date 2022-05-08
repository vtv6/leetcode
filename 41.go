package main

func firstMissingPositive(nums []int) int {
	for i, x := range nums {
		if x <= 0 || x > len(nums) {
			nums[i] = len(nums) + 1
		}
	}

	for i := range nums {
		curr := abs(nums[i])
		if curr <= len(nums) {
			nums[curr-1] = -abs(nums[curr-1])
		}
	}

	for i, x := range nums {
		if x > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
