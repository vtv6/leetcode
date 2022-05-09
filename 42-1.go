package main

func trap(height []int) int {
	stack := []int{0}
	result := 0
	for i := 1; i < len(height); i++ {
		if height[stack[len(stack)-1]] >= height[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
				currIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) > 0 {
					result += (min(height[i], height[stack[len(stack)-1]]) - height[currIndex]) * (i - stack[len(stack)-1] - 1)
				}
			}
			stack = append(stack, i)
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
