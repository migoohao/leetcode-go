package tree

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var result []float64
	var nextNodes []*TreeNode
	levelNods := []*TreeNode{root}
	sum, count := 0, 0
	for len(levelNods) != 0 {
		for _, node := range levelNods {
			sum += node.Val
			count++
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(count))
		levelNods, nextNodes = nextNodes, []*TreeNode{}
		sum = 0
		count = 0
	}
	return result
}
