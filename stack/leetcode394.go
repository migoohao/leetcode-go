package stack

import "strconv"

func decodeString(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		switch s[i] {
		case ']':
			tmp := make([]byte, 0)
			for v := stack[len(stack)-1]; v != '['; v = stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				tmp = append(tmp, v)
			}
			//pop '['
			stack = stack[:len(stack)-1]
			idx := len(stack) - 1
			for ; idx >= 0 && stack[idx] >= '0' && stack[idx] <= '9'; idx-- {
			}
			times, _ := strconv.Atoi(string(stack[idx+1:]))
			stack = stack[:idx+1]
			for i := 0; i < times; i++ {
				for l := len(tmp) - 1; l >= 0; l-- {
					stack = append(stack, tmp[l])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
