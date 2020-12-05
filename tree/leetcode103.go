package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	level := make([]*TreeNode, 0)
	if root != nil {
		level = append(level, root)
	}
	for len(level) > 0 {
		nextLevel := make([]*TreeNode, 0)
		sub := make([]int, 0)
		for len(level) > 0 {
			node := level[0]
			level = level[1:]
			sub = append(sub, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, sub)
		level = nextLevel
	}
	zflip(result)
	return result
}

func zflip(result [][]int) {
	for i := range result {
		if i%2 != 0 {
			reverse(result[i])
		}
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
