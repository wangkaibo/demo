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
	 	node = addNode(l[i], node)
	 }
	leetcode.ReverseList(&head)
	return
}

func addNode(val int, node *leetcode.ListNode) *leetcode.ListNode {
	node.Next = &leetcode.ListNode{
		Val: val,
		Next: nil,
	}
	return node;
}
