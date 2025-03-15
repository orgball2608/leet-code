package medium

func longestSubarray(nums []int, limit int) int {
	var minQueue, maxQueue []int
	left := 0
	maxLen := 0
	for right, num := range nums {
		for len(maxQueue) > 0 && nums[maxQueue[len(maxQueue)-1]] <= num {
			maxQueue = maxQueue[:len(maxQueue)-1]
		}

		maxQueue = append(maxQueue, right)

		for len(minQueue) > 0 && nums[minQueue[len(minQueue)-1]] >= num {
			minQueue = minQueue[:len(minQueue)-1]
		}

		minQueue = append(minQueue, right)

		for nums[maxQueue[0]]-nums[minQueue[0]] > limit {
			if maxQueue[0] == left {
				maxQueue = maxQueue[1:]
			}
			if minQueue[0] == left {
				minQueue = minQueue[1:]
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

// Time: o(N)
// Space: O(N)
