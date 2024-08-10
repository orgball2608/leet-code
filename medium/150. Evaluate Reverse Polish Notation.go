package medium

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	length := 0
	for _, val := range tokens {
		if i, err := strconv.Atoi(val); err == nil {
			stack = append(stack, i)
			length++
		} else {
			num1 := stack[length-1]
			num2 := stack[length-2]
			stack = stack[:length-2]
			length--
			switch val {
			case "+":
				stack = append(stack, num2+num1)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num2*num1)
			case "/":
				stack = append(stack, num2/num1)
			}
		}
	}

	return stack[len(stack)-1]
}

//Time: O(N)
//Space: O(N)
