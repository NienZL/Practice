package model

import (
	"strconv"
)

func Calc(input string) string {
	postfix := toPostfix(input)
	// fmt.Println(postfix)
	var stack []float64
	var n1, n2 float64
	for i := 0; i < len(postfix); i++ {
		n := len(stack)
		switch postfix[i] {
		case "+":
			n1 = stack[n-2]
			n2 = stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, n1+n2)
		case "-":
			n1 = stack[n-2]
			n2 = stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, n1-n2)
		case "*":
			n1 = stack[n-2]
			n2 = stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, n1*n2)
		case "/":
			n1 = stack[n-2]
			n2 = stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, n1/n2)
		default:
			nextInt, _ := strconv.ParseFloat(postfix[i], 64)
			stack = append(stack, nextInt)
		}
	}
	result := strconv.FormatFloat(stack[0], 'f', -1, 64)
	return result
}

// switch postfix[i] {
// case "+":
// 	stack = append(stack, n1+n2)
// case "-":
// 	stack = append(stack, n1-n2)
// case "*":
// 	stack = append(stack, n1*n2)
// case "/":
// 	stack = append(stack, n1/n2)
// }
