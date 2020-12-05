package dynamic

func jump(nums []int) int {
	steps := 0
	finish := len(nums) - 1
	for point, next := 0, 0; point < finish; steps++ {
		for i := nums[point]; i > 0; i-- {
			index := point + i
			if index >= finish {
				steps++
				return steps
			}
			nextNext := index + nums[index]
			if nextNext > next+nums[next] {
				next = index
				if nextNext >= finish {
					steps += 2
					return steps
				}
			}
		}
		point = next
	}
	return steps
}
