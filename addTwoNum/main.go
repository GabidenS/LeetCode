package main

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func main() {
	var l1, l2 *ListNode
	fmt.Println(l1, l2)
	pushBack(l1, 1)
	pushBack(l1, 2)
	pushBack(l1, 3)
	fmt.Println(l1, l2)
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var res *ListNode
// 	var ls1, ls2 int
// 	ls1 = listSize(l1)
// 	ls2 = listSize(l2)
// 	var tmp int
// 	if ls2 > ls1 {
// 		templ := l2
// 		for templ != nil {
// 			res.Val = (l1.Val+l2.Val)%10 + tmp
// 			tmp = 0
// 			if (l1.Val+l2.Val)/10 == 1 {
// 				tmp = 1
// 			}

// 		}
// 	}
// 	return res
// }

func listSize(l *ListNode) int {
	count := 0
	n := l
	for n != nil {
		n = n.Next
		count++
	}
	return count
}

func pushBack(l *ListNode, d int) *ListNode {
	if l == nil {
		l.Val = d
	} else {
		n := l
		for n != nil {
			n = n.Next
		}
		n.Val = d
		l = n
	}
	return l

}
