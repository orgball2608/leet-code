package easy

func missingNumber(nums []int) int {
	expect := 0
	actual := 0
	for i := 1; i <= len(nums); i++ {
		expect += i
		actual += nums[i-1]
	}
	return expect - actual
}

// Time: O(N)
//Space: O(1)
