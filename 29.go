package main

import "math"

func divide(dividend int, divisor int) int {
	MAX := 2147483647
	if divisor == -1 && dividend <= -MAX-1 {
		return MAX
	} else if divisor == -1 {
		return -dividend
	} else if divisor == 1 {
		return dividend
	}

	positive := true

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		positive = false
	}

	f_dividend := math.Abs(float64(dividend))
	f_divisor := math.Abs(float64(divisor))

	result := 0
	for f_dividend >= f_divisor && result < MAX {
		result++
		f_dividend = f_dividend - f_divisor
	}

	if positive {
		return result
	}
	return -result
}
