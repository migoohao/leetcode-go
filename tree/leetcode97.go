package tree

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inOrder(root, &result)
	return result
}

func recursiveInOrder(root *TreeNode, result *[]int) {
	if root != nil {
		recursiveInOrder(root.Left, result)
		*result = append(*result, root.Val)
		recursiveInOrder(root.Right, result)
	}
}

func inOrder(root *TreeNode, result *[]int) {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for node := root; node != nil; node = node.Left {
			stack = append(stack, node)
		}
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		*result = append(*result, node.Val)
		root = node.Right
	}
}
