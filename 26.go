package main

func removeDuplicates(nums []int) int {
	pointer := 0
	for pointer < len(nums)-1 {
		if nums[pointer] == nums[pointer+1] {
			nums = append(nums[:pointer+1], nums[pointer+2:]...)
		} else {
			pointer++
		}
	}

	return len(nums)
}
