package main

func main() {
	search([]int{5, 1, 3}, 1)
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {

	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	middle := (left + right) / 2

	if nums[left] < nums[middle] {
		if nums[left] <= target && nums[middle] >= target {
			return binarySearch(nums, left, middle, target)
		} else {
			return binarySearch(nums, middle+1, right, target)
		}
	} else {
		if nums[middle+1] <= target && nums[right] >= target {
			return binarySearch(nums, middle+1, right, target)
		} else {
			return binarySearch(nums, left, middle, target)
		}
	}
}
