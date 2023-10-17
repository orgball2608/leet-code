func countPairs(nums []int, target int) int {
    sort.Ints(nums)
    left,right,sum := 0, len(nums)-1, 0
    length := 0
    for left < right {
        sum = nums[left] + nums[right]
        if(sum < target) {
            length += right-left;
            left++
        } else {
            right--
        }
    }
    return length
}