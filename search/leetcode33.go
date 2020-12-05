package search

func searchIII(nums []int, target int) int {
	start, end := 0, len(nums)-1
	mid := -1
	for mid = (start + end) / 2; start <= end; mid = (start + end) / 2 {
		if nums[mid] == target {
			return mid
		}
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}
		if nums[start] < nums[mid] {
			if nums[start] < target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
