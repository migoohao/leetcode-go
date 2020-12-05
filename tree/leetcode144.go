package tree

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	preOrder(root, &result)
	return result
}

func recursivePre(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		recursivePre(root.Left, result)
		recursivePre(root.Right, result)
	}
}

func preOrder(root *TreeNode, result *[]int) {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for node := root; node != nil; node = node.Left {
			*result = append(*result, node.Val)
			stack = append(stack, node)
		}
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = node.Right
	}
}
