package tree

import (
	"strconv"
	"strings"
)

var (
	pathElements []string
	result       []string
)

const (
	sep = "->"
)

func binaryTreePaths(root *TreeNode) []string {
	pathElements = []string{}
	result = []string{}
	if root != nil {
		dfs(root)
	}
	return result
}

func dfs(root *TreeNode) {
	pathElements = append(pathElements, strconv.Itoa(root.Val))
	currentIndex := len(pathElements)
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	if root.Left == nil && root.Right == nil {
		result = append(result, strings.Join(pathElements, sep))
	}
	pathElements = pathElements[:currentIndex-1]
}
