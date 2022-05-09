package main

func trap(height []int) int {
	maxLefts := make([]int, len(height))
	maxLefts[0] = height[0]
	for i := 1; i < len(height); i++ {
		if maxLefts[i-1] < height[i] {
			maxLefts[i] = height[i]
		} else {
			maxLefts[i] = maxLefts[i-1]
		}
	}

	maxRights := make([]int, len(height))
	maxRights[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if maxRights[i+1] < height[i] {
			maxRights[i] = height[i]
		} else {
			maxRights[i] = maxRights[i+1]
		}
	}

	sum := 0
	for i := 1; i < len(height)-1; i++ {
		sum += min(maxLefts[i], maxRights[i]) - height[i]
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
