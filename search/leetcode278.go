package search

func firstBadVersion(n int) int {
	start, end := 1, n+1
	for mid := (start + end) / 2; start < end; mid = (start + end) / 2 {
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if isBadVersion(start) {
		return start
	} else {
		return start + 1
	}
}

func isBadVersion(version int) bool {
	return true
}
