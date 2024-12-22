package easy

func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for i, val := range nums {
		if val != 0 {
			nums[nonZeroIndex] = val
			nonZeroIndex++
		}

		if i >= nonZeroIndex {
			nums[i] = 0
		}
	}
}
