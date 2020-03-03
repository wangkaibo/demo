package leetcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	return head
}