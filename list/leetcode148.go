package list

func sortList(head *ListNode) *ListNode {
	result, _ := mergeSort(head, nil)
	return result
}

func mergeSort(head, tail *ListNode) (front, rear *ListNode) {
	if head == tail || head.Next == tail {
		return head, head
	}
	slow := head
	fast := head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != tail {
		fast = fast.Next
	}
	front1, rear1 := mergeSort(head, slow)
	front2, rear2 := mergeSort(slow, tail)
	return merge2Links(front1, rear1.Next, front2, rear2.Next)
}

func merge2Links(l1, end1, l2, end2 *ListNode) (front, rear *ListNode) {
	fake := &ListNode{}
	rear = fake
	for l1 != end1 && l2 != end2 {
		if l1.Val <= l2.Val {
			rear.Next = l1
			l1 = l1.Next
		} else {
			rear.Next = l2
			l2 = l2.Next
		}
		rear = rear.Next
	}
	for l1 != end1 {
		rear.Next = l1
		l1 = l1.Next
		rear = rear.Next
	}
	for l2 != end2 {
		rear.Next = l2
		l2 = l2.Next
		rear = rear.Next
	}
	rear.Next = nil
	return fake.Next, rear
}
