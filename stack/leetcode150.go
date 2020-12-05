package stack

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, s := range tokens {
		idx := len(stack) - 2
		value := 0
		switch s {
		case "+":
			value = stack[idx] + stack[idx+1]
			stack = stack[:idx]
		case "-":
			value = stack[idx] - stack[idx+1]
			stack = stack[:idx]
		case "*":
			value = stack[idx] * stack[idx+1]
			stack = stack[:idx]
		case "/":
			value = stack[idx] / stack[idx+1]
			stack = stack[:idx]
		default:
			value, _ = strconv.Atoi(s)
		}
		stack = append(stack, value)
	}
	return stack[0]
}
