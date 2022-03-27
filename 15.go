package main

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1

		for l < r && l < len(nums) && r >= 0 {
			if nums[l]+nums[r] < -nums[i] {
				l++
			} else if nums[l]+nums[r] > -nums[i] {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})

				for l < len(nums)-1 && nums[l+1] == nums[l] {
					l++
				}
				for r > 0 && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}
