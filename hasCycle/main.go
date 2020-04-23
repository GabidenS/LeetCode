package main

func main() {

}

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var mCycle = make(map[*ListNode]int)
	if head == nil {
		return false
	}
	for head != nil {
		if _, ok := mCycle[head]; ok {
			return true
		} else {
			mCycle[head] = 1
			head = head.Next
		}

	}
	return false
}
