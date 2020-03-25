package leetcode

func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	if head == nil || head.Next == nil {
		return false
	}
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast && fast != nil {
			return true
		}
	}

	return false
}