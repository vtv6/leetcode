package main

func wiggleMaxLength(nums []int) int {
	result := 1
	prevDiff := 0
	diff := 0

	for i := 1; i < len(nums); i++ {
		diff = nums[i] - nums[i-1]
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			result++
			prevDiff = diff
		}
	}

	return result
}
