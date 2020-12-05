package search

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for mid := (start + end) / 2; start < end; mid = (start + end) / 2 {
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
