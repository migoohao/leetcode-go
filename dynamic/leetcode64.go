package dynamic

import "math"

func minPathSum(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	help := make([]int, len(grid[0])+1)
	for i, value := range grid[0] {
		help[i+1] = help[i] + value
	}
	help[0] = math.MaxInt32
	for k := 1; k < len(grid); k++ {
		for i, value := range grid[k] {
			help[i+1] = min(help[i], help[i+1]) + value
		}
	}
	return help[len(help)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
