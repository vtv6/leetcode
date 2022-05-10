package main

func combinationSum3(k int, n int) [][]int {
	return calc(k, n, 1)
}

func calc(k int, n int, l int) [][]int {
	if k == 1 {
		if n >= l && n <= 9 {
			return [][]int{[]int{n}}
		} else {
			return [][]int{}
		}
	}
	s := [][]int{}

	for i := l; i <= 9; i++ {
		x := calc(k-1, n-i, i+1)
		for _, j := range x {
			s = append(s, append(append([]int{}, i), j...))
		}
	}
	return s
}
