package medium

func numSubarraysWithSum(nums []int, goal int) int {
	left1, left2, sum1, sum2, count := 0, 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum1 += nums[i]
		sum2 += nums[i]
		for left1 <= i && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}

		for left2 <= i && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		count += left2 - left1
	}

	return count
}

// Time: O(N)
// Space: O(1)
