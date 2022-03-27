package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		temp := threeSum(nums[i+1:], target-nums[i])

		for _, triple := range temp {
			result = append(result, append([]int{nums[i]}, triple...))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

func threeSum(nums []int, target int) [][]int {
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
			if nums[l]+nums[r]+nums[i] < target {
				l++
			} else if nums[l]+nums[r]+nums[i] > target {
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
