func threeSumClosest(nums []int, target int) int {
    if len(nums) == 3 {
        return nums[0] + nums[1] + nums[2]
    }
    sort.Ints(nums)
    var ans int
    minDiff := math.MaxInt
    for i := 0; i < len(nums) - 2; i++ {
        left, right := i+1, len(nums)-1
        for left < right {
            diff := target - nums[i] - nums[left] - nums[right]
            if diff == 0 {
                return target
            } else if abs(diff) < minDiff {
                minDiff = abs(diff)
                ans = nums[i] + nums[left] + nums[right]
            }
            if diff > 0 {
                left++
            } else {
                right--
            }
        }
    }
    return ans
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
