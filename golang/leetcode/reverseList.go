package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var current *ListNode
	var temp *ListNode
	for {
		current = head
		temp = current
		current.Next = prev
		prev = current
		current = temp
		break
	}

	return prev
}
