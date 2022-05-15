package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	s, _ := calc(root)
	return s
}

func calc(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ls, l := calc(root.Left)
	rs, r := calc(root.Right)

	if l == r && l == 0 {
		return root.Val, 1
	}

	if l > r {
		return ls, l + 1
	} else if l < r {
		return rs, r + 1
	} else {
		return ls + rs, l + 1
	}
}
