package main

import (
	"math"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var lm, rm, ld, rd int

	if root == nil {
		return 0
	}
	lm = maxDepth(root.Left)
	rm = maxDepth(root.Right)

	ld = diameterOfBinaryTree(root.Left)
	rd = diameterOfBinaryTree(root.Right)
	return max(lm+rm, max(rd, ld))

}

func maxDepth(root *TreeNode) int {
	// var ld, rd int
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
	sort.Ints()
}
