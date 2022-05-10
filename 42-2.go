package main

func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	for l < r {
		if height[l] < height[r] {
			if leftMax < height[l] {
				leftMax = height[l]
			} else {
				result += leftMax - height[l]
				l++
			}
		} else {
			if rightMax < height[r] {
				rightMax = height[r]
			} else {
				result += rightMax - height[r]
				r--
			}
		}
	}

	return result
}
