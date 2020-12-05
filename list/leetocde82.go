package list

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fakeHeader := &ListNode{Next: head}
	pre := fakeHeader
	cur := head.Next
	for pre.Next != nil && cur != nil {
		if pre.Next.Val == cur.Val {
			for cur != nil && pre.Next.Val == cur.Val {
				cur = cur.Next
			}
			pre.Next = cur
			if cur != nil {
				cur = cur.Next
			}
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	return fakeHeader.Next
}
