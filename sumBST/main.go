package main

import "fmt"

func main() {

}

////Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		fmt.Println("this is nil")
		return 0
	}
	if root.Val >= L && root.Val <= R {
		fmt.Println(root.Val)
		return rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R) + root.Val
	} else {
		fmt.Println("this is else state:", root.Val)
		rangeSumBST(root.Right, L, R)
		rangeSumBST(root.Left, L, R)

	}
	fmt.Println("this is 0")
	return 0
}
