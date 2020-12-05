package dynamic

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	help := make([]int, n+1)
	for i := 1; i <= n; i++ {
		help[i] = 1
	}
	for r := 1; r < m; r++ {
		for i := 1; i <= n; i++ {
			help[i] += help[i-1]
		}
	}
	return help[n]
}
