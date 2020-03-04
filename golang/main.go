package main

import (
	"demo/golang/leetcode"
)

func main() {
	 l := []int{1,2,3,4,5}
	head := leetcode.ListNode{
		Val: l[0],
		Next: nil,
	}
	node := &head;
	for i := 1; i < len(l); i++ {
		node.Next = &leetcode.ListNode{
			Val: l[i],
			Next: nil,
		}
		node = node.Next
	 }
	leetcode.ReverseList(&head)
	return
}
