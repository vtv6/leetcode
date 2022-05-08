package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}

	sort.Ints(candidates)

	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		x := candidates[i]
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

	for i := index + 1; i < len(candidates); i++ {

		if (s+candidates[i]) > target || (i > index+1 && candidates[i] == candidates[i-1]) {
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
