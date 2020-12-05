package dynamic

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	help := make([]int, len(s))
	maxEnd := -1
	for i := 0; i < len(s); i++ {
		for k := 0; k < 2; k++ {
			start, end := findPalindrome(s, i, i+k)
			if start <= end && end > maxEnd {
				if start <= 0 {
					setHelper(help, start, end, 0)
				} else {
					setHelper(help, start, end, help[start-1]+1)
				}
				maxEnd = end
			}
		}
	}
	return help[len(s)-1]
}

func setHelper(help []int, from, to, value int) {
	for i := from; i <= to; i++ {
		help[i] = value
	}
}

func findPalindrome(s string, midL, midR int) (int, int) {
	for midL >= 0 && midR < len(s) && s[midL] == s[midR] {
		midL--
		midR++
	}
	midL++
	midR--
	return midL, midR
}
