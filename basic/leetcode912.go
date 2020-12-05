package basic

func sortArray(nums []int) []int {
	//quickSort(nums)
	mergeSort(nums)
	return nums
}

func quickSort(nums []int) {
	if len(nums) > 1 {
		mid := partition(nums)
		quickSort(nums[:mid])
		quickSort(nums[mid+1:])
	}
}

func partition(nums []int) int {
	last := len(nums) - 1
	p := nums[last]
	i, j := 0, 0
	for ; j < last; j++ {
		if nums[j] <= p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[last] = nums[last], nums[i]
	return i
}

func mergeSort(nums []int) {
	if len(nums) > 1 {
		mid := len(nums) / 2
		mergeSort(nums[:mid])
		mergeSort(nums[mid:])
		first := 0
		second := mid
		assist := make([]int, 0)
		for first < mid && second < len(nums) {
			if nums[first] < nums[second] {
				assist = append(assist, nums[first])
				first++
			} else {
				assist = append(assist, nums[second])
				second++
			}
		}
		assist = append(assist, nums[first:mid]...)
		assist = append(assist, nums[second:]...)
		copy(nums, assist)
	}
}
