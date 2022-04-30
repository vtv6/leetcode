package main

func main() {
	combinationSum([]int{2, 3, 1}, 9)
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	for i, x := range candidates {
		if x > target {
			continue
		} else if x == target {
			result = append(result, []int{x})
			continue
		}

		a := backTracking(candidates, target, i, []int{candidates[i]})
		result = append(result, a...)
	}

	return result
}

func backTracking(candidates []int, target int, index int, sn []int) [][]int {
	result := [][]int{}

	// wrong answer if remove this. don't know why
	startNums := make([]int, len(sn))
	copy(startNums, sn)

	s := 0
	for _, x := range startNums {
		s += x
	}

	for i := index; i < len(candidates); i++ {
		if (s + candidates[i]) > target {
			continue
		} else if (s + candidates[i]) == target {
			result = append(result, append(startNums, candidates[i]))
			continue
		}
		a := backTracking(candidates, target, i, append(startNums, candidates[i]))
		result = append(result, a...)
	}

	return result
}
