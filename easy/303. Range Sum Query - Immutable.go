package easy

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left-1 < 0 {
		return this.Nums[right]
	}

	return this.Nums[right] - this.Nums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

// Time; O(1)
// Space: O(N)
