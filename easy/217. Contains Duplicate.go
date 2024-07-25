package easy

func containsDuplicate(nums []int) bool {
	a := make(map[int]bool)
	for _, val := range nums {
		if _, ok := a[val]; ok {
			return true
		}
		a[val] = true
	}
	return false
}
