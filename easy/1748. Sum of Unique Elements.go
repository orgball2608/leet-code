package easy

func sumOfUnique(nums []int) int {
	hashMap := make(map[int]int)
	for _, val := range nums {
		hashMap[val] += 1
	}

	sum := 0
	for key, val := range hashMap {
		if val == 1 {
			sum += key
		}
	}
	return sum
}

// Time: O(N)
// Space: O(N)
