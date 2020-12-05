package search

func findMinII(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		front := nums[start]
		for start <= end && front == nums[start] {
			start++
		}
		start--
		rear := nums[end]
		for start <= end && rear == nums[end] {
			end--
		}
		end++
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
