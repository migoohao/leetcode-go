package dynamic

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) <= 0 {
		return 0
	}
	help := make([]int, len(obstacleGrid[0])+1)
	for r := 0; r < len(obstacleGrid); r++ {
		for c := 0; c < len(obstacleGrid[r]); c++ {
			if obstacleGrid[r][c] == 1 {
				help[c+1] = 0
			} else if r == 0 && c == 0 {
				help[c+1] = 1
			} else {
				help[c+1] += help[c]
			}
		}
	}
	return help[len(help)-1]
}
