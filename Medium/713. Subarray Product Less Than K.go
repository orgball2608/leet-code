func numSubarrayProductLessThanK(nums []int, k int) int {
    n := len(nums)
    l,acc,res:=0,1,0
    for r:=0;r<n;r++ {
        acc*=nums[r]
        for acc >= k && l <= r {
            acc /= nums[l]
            l++
        }
        res+=r-l+1
    }
    return res
}