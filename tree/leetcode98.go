package tree

func isValidBST(root *TreeNode) bool {
	path := travelList(root)
	return increasing(path)
}

func travelList(root *TreeNode) []int {
	path := make([]int, 0)
	middleOrder(root, &path)
	return path
}

func middleOrder(root *TreeNode, path *[]int) {
	if root != nil {
		middleOrder(root.Left, path)
		*path = append(*path, root.Val)
		middleOrder(root.Right, path)
	}
}

func increasing(path []int) bool {
	if len(path) <= 1 {
		return true
	}
	pre := path[0]
	for i := 1; i < len(path); i++ {
		if pre >= path[i] {
			return false
		}
		pre = path[i]
	}
	return true
}
