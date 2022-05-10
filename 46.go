package main

func main() {
	permute([]int{1, 2, 3})
}
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	s := [][]int{}

	for i, x := range nums {
		cop := make([]int, len(nums))
		copy(cop, nums)
		arr := permute(append(cop[:i], cop[i+1:]...))
		for j := range arr {
			s = append(s, append(arr[j], x))
		}
	}

	return s
}
