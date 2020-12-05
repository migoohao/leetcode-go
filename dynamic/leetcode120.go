package dynamic

func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	help := make([]int, len(triangle))
	for _, arr := range triangle {
		for i := len(arr) - 1; i >= 0; i-- {
			if i == 0 {
				help[i] += arr[i]
			} else if i == len(arr)-1 {
				help[i] = help[i-1] + arr[i]
			} else {
				help[i] = min(help[i], help[i-1]) + arr[i]
			}
		}
	}
	mininum := math.MaxInt32
	for _, v := range help {
		mininum = min(mininum, v)
	}
	return mininum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
