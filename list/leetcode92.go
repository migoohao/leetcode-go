package list

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	result := &ListNode{Next: head}
	pre := result
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	reverseLastNode := pre.Next
	var reverseTopNode *ListNode
	node := pre.Next
	for i := 0; i < n-m+1; i++ {
		tmp := node.Next
		node.Next = reverseTopNode
		reverseTopNode = node
		node = tmp
	}
	pre.Next = reverseTopNode
	reverseLastNode.Next = node
	return result.Next
}
