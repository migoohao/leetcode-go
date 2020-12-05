package dynamic

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	p := 2
	pp := 1
	cur := 0
	for i := 3; i <= n; i++ {
		cur = p + pp
		pp = p
		p = cur
	}
	return cur
}
