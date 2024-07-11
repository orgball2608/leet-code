package easy

// 100% Space
func merge(nums1 []int, m int, nums2 []int, n int) {
	var pt1, pt2, pt3 int = m - 1, n - 1, m + n - 1
	for ; pt1 >= 0 && pt2 >= 0; pt3-- {
		if nums1[pt1] <= nums2[pt2] {
			nums1[pt3] = nums2[pt2]
			pt2--
		} else {
			nums1[pt3] = nums1[pt1]
			pt1--
		}
	}

	if pt2 >= 0 {
		copy(nums1[:pt3+1], nums2[:pt2+1])
	}
}

//Time: O(N+M)
//Space: O(1)

// 100% Time
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	i, j := 1, 1
//	result := make([]int, m+n)
//	idx := 0
//	for i <= m || j <= n {
//		if j == n+1 || (i <= m && nums1[i-1] <= nums2[j-1]) {
//			result[idx] = nums1[i-1]
//			i++
//		} else {
//			result[idx] = nums2[j-1]
//			j++
//		}
//		idx++
//	}
//
//	for i, value := range result {
//		nums1[i] = value
//	}
//}

//Time: O(2N+2M)
//Space: O(N+M)
