package main

func permuteUnique(nums []int) [][]int {
	m := make(map[int][]int)

	for _, num := range nums {
		m[num] = append(m[num], num)
	}

	return permute2(m)
}

func permute2(nums map[int][]int) [][]int {
	s := [][]int{}
	for k, v := range nums {
		if len(v) == 0 {
			continue
		}
		nums[k] = nums[k][1:]
		arr := permute2(nums)
		for i := range arr {
			s = append(s, append(arr[i], k))
		}
		if len(arr) == 0 {
			s = append(s, []int{k})
		}
		nums[k] = append(nums[k], k)
	}

	return s
}
