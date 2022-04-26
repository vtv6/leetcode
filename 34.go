package main

func searchRange(nums []int, target int) []int {
	var startingPosition, endingPosition int
	startingPosition = bs(nums, target, 0, len(nums)-1, true)
	if startingPosition == -1 {
		endingPosition = -1
	} else {
		endingPosition = bs(nums, target, startingPosition, len(nums)-1, false)
	}
	return []int{startingPosition, endingPosition}

}

func bs(nums []int, target, l, r int, startingPosition bool) int {
	if r < l {
		return -1
	}
	if l == r {
		if nums[l] == target {
			return l
		} else {
			return -1
		}
	}

	m := (l + r) / 2

	if nums[m] > target {
		return bs(nums, target, l, m, startingPosition)
	} else if nums[m] < target {
		return bs(nums, target, m+1, r, startingPosition)
	} else {
		if startingPosition {
			return min(bs(nums, target, l, m-1, startingPosition), m)
		} else {
			return max(bs(nums, target, m+1, r, startingPosition), m)
		}
	}
}

func min(x, y int) int {
	if x < y && x != -1 {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y && x != -1 {
		return x
	} else {
		return y
	}
}
