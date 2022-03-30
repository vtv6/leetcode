package main

func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i-1]
			index++
		}
	}
	nums[index] = nums[len(nums)-1]
	return index + 1
}
