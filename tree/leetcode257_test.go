package tree

import (
	"testing"
)

var (
	five  = TreeNode{5, nil, nil}
	three = TreeNode{3, nil, nil}
	two   = TreeNode{2, nil, &five}
	one   = TreeNode{1, &two, &three}
)

func Test_binaryTreePaths(t *testing.T) {
	result := binaryTreePaths(&one)
	if len(result) != 2 || result[0] != "1->2->5" || result[1] != "1->3" {
		t.Error("result fail:", result)
	}
}

func Benchmark_binaryTreePaths(b *testing.B) {
	binaryTreePaths(&one)
}
