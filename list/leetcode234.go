package list

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := reverse(slow.Next)
	for l2 != nil {
		if head.Val != l2.Val {
			return false
		}
		head = head.Next
		l2 = l2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = node
		node = head
		head = tmp
	}
	return node
}
