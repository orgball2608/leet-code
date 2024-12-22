package easy

func findDifference(nums1 []int, nums2 []int) [][]int {
	hashMap1 := make(map[int]struct{}, len(nums1))
	for _, val := range nums1 {
		hashMap1[val] = struct{}{}
	}

	hashMap2 := make(map[int]struct{}, len(nums2))
	for _, val := range nums2 {
		hashMap2[val] = struct{}{}
	}

	for _, val := range nums1 {
		if _, ok := hashMap2[val]; ok {
			delete(hashMap1, val)
			delete(hashMap2, val)
		}
	}

	result1 := make([]int, 0, len(hashMap1))
	for key, _ := range hashMap1 {
		result1 = append(result1, key)
	}

	result2 := make([]int, 0, len(hashMap2))
	for key, _ := range hashMap2 {
		result2 = append(result2, key)
	}

	return [][]int{result1, result2}
}

// Time: O(2N+M+P+Q)
// Space: O(N+M+P+Q)
