package list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	duplica := make(map[*Node]*Node)
	for node := head; node != nil; node = node.Next {
		duplica[node] = &Node{Val: node.Val}
	}
	for node := head; node != nil; node = node.Next {
		if node.Next != nil {
			duplica[node].Next = duplica[node.Next]
		}
		if node.Random != nil {
			duplica[node].Random = duplica[node.Random]
		}
	}
	return duplica[head]
}
