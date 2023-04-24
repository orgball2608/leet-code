func findMaxAverage(nums []int, k int) float64 {
    maxSum := math.MinInt
    Wdsum := 0
    for i:=0 ; i< len(nums); i++ {
        Wdsum += nums[i]
        if i >= k -1 {
            if i !=k -1 {
                Wdsum -= nums[i-k]
            }
            if Wdsum > maxSum {
                maxSum = Wdsum
            }
        }
    }
    // var result float64 
    return float64(maxSum)/float64(k)
}