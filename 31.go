package main

import "sort"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i >= 0 {
		min := nums[i+1]
		minIndex := i + 1
		for j := i + 2; j < len(nums); j++ {
			if min > nums[j] && nums[j] > nums[i] {
				minIndex = j
				min = nums[j]
			}
		}

		temp := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = temp
	}

	sort.Ints(nums[i+1:])

}
