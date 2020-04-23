package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	l := head
	var res int
	if head == nil {
		return 0
	} else {

		for l != nil {
			// s:=strconv.Itoa(l.Val)
			res = res*2 + l.Val
			l = l.Next
		}
	}
	return res
}
