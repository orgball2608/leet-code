package easy

func removeDuplicates(s string) string {
	stack := []rune{}
	for _, val := range s {
		if len(stack) > 0 && stack[len(stack)-1] == val {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, val)
		}
	}

	return string(stack)
}

//Time: O(N)
//Space: O(1)
