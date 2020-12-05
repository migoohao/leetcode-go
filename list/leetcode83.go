package list

func deleteDuplicates(head *ListNode) *ListNode {
	result := head
	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return result
}
