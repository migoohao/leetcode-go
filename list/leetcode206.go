package list

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	result = nil
	for head != nil {
		tmp := head.Next
		head.Next = result
		result = head
		head = tmp
	}
	return result
}
