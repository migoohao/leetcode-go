package tree

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	postOrder(root, &result)
	return result
}

func recursivePostOrder(root *TreeNode, result *[]int) {
	if root != nil {
		recursivePostOrder(root.Left, result)
		recursivePostOrder(root.Right, result)
		*result = append(*result, root.Val)
	}
}

func postOrder(root *TreeNode, result *[]int) {
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) > 0 {
		for node := root; node != nil; node = node.Left {
			stack = append(stack, node)
		}
		parent := stack[len(stack)-1]
		if parent.Right == nil || parent.Right == lastVisit {
			*result = append(*result, parent.Val)
			stack = stack[:len(stack)-1]
			lastVisit = parent
			root = nil
		} else {
			root = parent.Right
		}
	}
}
