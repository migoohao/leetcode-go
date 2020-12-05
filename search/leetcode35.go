package search

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)
	for mid := (start + end) / 2; start < end; mid = (start + end) / 2 {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(nums) || nums[start] > target {
		return start
	}
	return start + 1
}
