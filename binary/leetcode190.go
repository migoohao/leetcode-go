package binary

func reverseBits(num uint32) uint32 {
	var result uint32
	for pow := 31; num != 0; pow-- {
		result |= (num & 1) << pow
		num >>= 1
	}
	return result
}
