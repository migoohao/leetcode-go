package search

func search(nums []int, target int) bool {
	if len(nums) <= 0 {
		return false
	}
	start, end := 0, len(nums)-1
	for start <= end {
		head := nums[start]
		if head == target {
			return true
		}
		for start <= end && nums[start] == head {
			start++
		}
		start--
		tail := nums[end]
		if tail == target {
			return true
		}
		for start <= end && nums[end] == tail {
			end--
		}
		end++
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if head <= nums[mid] {
			if head < target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] <= tail {
			if nums[mid] < target && target < tail {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
