package binary

func singleNumberII(nums []int) int {
	y, x := 0, 0
	for _, v := range nums {
		old := y
		y = (v ^ y) & (^x)
		x = (x & (^old) & (^v)) | ((^x) & old & v)
	}
	return y
}
