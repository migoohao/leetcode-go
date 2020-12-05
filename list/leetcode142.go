package list

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head.Next
	fast := head.Next.Next
	var collision *ListNode
	for fast != nil && fast.Next != nil {
		if slow == fast {
			collision = fast
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if collision == nil {
		return nil
	}
	for head != collision {
		head = head.Next
		collision = collision.Next
	}
	return collision
}
