package binary

func hammingWeight(num uint32) int {
	weight := 0
	for ; num != 0; weight++ {
		num &= num - 1
	}
	return weight
}
