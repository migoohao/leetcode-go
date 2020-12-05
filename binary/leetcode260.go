package binary

func singleNumberIII(nums []int) []int {
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	diff := (sum & (sum - 1)) ^ sum
	result := []int{0, 0}
	for _, v := range nums {
		if (v & diff) == 0 {
			result[0] ^= v
		} else {
			result[1] ^= v
		}
	}
	return result
}
