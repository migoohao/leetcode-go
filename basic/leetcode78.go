package basic

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backTrack(nums, 0, list, &result)
	return result
}

func backTrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for ; pos < len(nums); pos++ {
		ansLen := len(list)
		list = append(list, nums[pos])
		backTrack(nums, pos+1, list, result)
		list = list[0:ansLen]
	}
}
