package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	closest := math.Inf(1)

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
outer:
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == target {
				result = target
				break outer
			} else if nums[i]+nums[l]+nums[r] > target {
				if closest > math.Abs(float64(nums[i]+nums[l]+nums[r]-target)) {
					closest = math.Abs(float64(nums[i] + nums[l] + nums[r] - target))
					result = nums[i] + nums[l] + nums[r]
				}

				r--
			} else {
				if closest > math.Abs(float64(nums[i]+nums[l]+nums[r]-target)) {
					closest = math.Abs(float64(nums[i] + nums[l] + nums[r] - target))
					result = nums[i] + nums[l] + nums[r]
				}

				l++
			}
		}
	}

	return result
}
