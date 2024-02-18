func containsNearbyDuplicate(nums []int, k int) bool {
    data := make(map[int]int)
    for i, v := range nums {
        index, ok := data[v]
        if ok && (i-index) <= k {
            return true
        }
        data[v]= i
    }
    return false
}
