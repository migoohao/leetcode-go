package stack

func largestRectangleArea(heights []int) int {
	areaMax := 0
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		cur := 0
		if i < len(heights) {
			cur = heights[i]
		}
		for len(stack) > 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			w := i
			if len(stack) > 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			areaMax = max(areaMax, w*h)
		}
		stack = append(stack, i)
	}
	return areaMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
