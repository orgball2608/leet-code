package easy

func findComplement(num int) int {
	temp := num
	count := 0
	for temp > 0 {
		num = swap(num, count)
		count++
		temp = temp >> 1
	}
	return num
}

func swap(u int, k int) int {
	return u ^ (1 << k)
}

// O(logN)
// Space: O(1)

func findComplement2(num int) int {
	temp := num
	length := 0
	for temp > 0 {
		length++
		temp = temp >> 1
	}
	return ((1 << length) - 1) ^ num
}

// O(logN)
// Space: O(1)
