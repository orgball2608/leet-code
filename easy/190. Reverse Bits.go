package easy

func reverseBits(num uint32) uint32 {
	left, right := 0, 31
	for left < right {
		i := getBit(num, left)
		j := getBit(num, right)
		if i != j {
			num = swap(num, left)
			num = swap(num, right)
		}
		left++
		right--
	}
	return num
}

func getBit(n uint32, k int) uint32 {
	return (n >> k) & 1
}

func swap(n uint32, k int) uint32 {
	return n ^ (1 << k)
}

// Time: O(1)
// Space: O(1)
