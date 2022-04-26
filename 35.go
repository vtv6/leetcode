package main

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
