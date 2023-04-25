func minSubArrayLen(target int, nums []int) int {
    start := 0
    sum := 0
    ml := len(nums) + 1
    for i, value := range nums {
        sum += value 
        for sum >= target {
            // if ml > (i - start +1) {
            //     ml = (i-start +1)
            // }
            ml = min(ml, (i - start +1))
            sum -= nums[start]
            start++
        }
    }
    if ml == len(nums) + 1 {
        return 0 
    }
    return ml 
}

func min(a,b int) int{
    if a <b {
        return a
    }
    return b
}