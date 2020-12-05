package list

func partition(head *ListNode, x int) *ListNode {
	result := &ListNode{Val: x - 1, Next: head}
	node := result
	for node.Next != nil && node.Next.Val < x {
		node = node.Next
	}
	preTail := node
	for node.Next != nil {
		for node.Next != nil && node.Next.Val >= x {
			node = node.Next
		}
		if node.Next == nil {
			break
		}
		tmp := node.Next
		node.Next = node.Next.Next
		tmp.Next = preTail.Next
		preTail.Next = tmp
		preTail = tmp
	}
	return result.Next
}
