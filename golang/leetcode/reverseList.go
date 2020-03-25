package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNode(s []int) *ListNode {
	head := ListNode{
		Val: s[0],
		Next: nil,
	}
	current := &head
	for i := 1; i< len(s) - 1; i++ {
		current.Next = &ListNode{
			Val: s[i],
			Next: nil,
		}
		current = current.Next
	}
	return &head
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
