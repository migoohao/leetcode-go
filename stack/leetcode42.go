package stack

func trap(height []int) int {
	stack := make([]int, 0)
	water := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				h := min(height[stack[len(stack)-1]], height[i]) - height[pop]
				w := i - stack[len(stack)-1] - 1
				water += h * w
			}
		}
		stack = append(stack, i)
	}
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
