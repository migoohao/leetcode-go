package dp

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	if dungeon == nil {
		return 1
	}
	m := len(dungeon)
	n := len(dungeon[0])
	if m <= 0 || n <= 0 {
		return 1
	}
	dp := initDp(m+1, n+1)
	dp[m][n-1], dp[m-1][n] = 1, 1
	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			dp[row][col] = max(min(dp[row+1][col], dp[row][col+1])-dungeon[row][col], 1)
		}
	}
	return dp[0][0]
}

func initDp(m, n int) [][]int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		for k, _ := range dp[i] {
			dp[i][k] = math.MaxInt32
		}
	}
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
