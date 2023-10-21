func fourSum(nums []int, target int) [][]int {
	var results [][]int
    if len(nums) < 4 {
        return results
    }
	sort.Ints(nums)
	for i := 0; i < len(nums) - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
            if j > i + 1 && nums[j] == nums[j-1] {
                continue
            }
			left, right := j + 1, len(nums)-1
			targetSum := target - nums[i] - nums[j]
			for left < right {
				sum := nums[left] + nums[right]

				if sum == targetSum {
					results = append(results, []int{nums[i],nums[j], nums[left], nums[right]})
					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < targetSum {
					left++
				} else {
					right--
				}
			}

		}
	}
	return results
}