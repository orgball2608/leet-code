func sortedSquares(nums []int) []int {
   res := make([]int, len(nums))
   l,r := 0, len(nums) - 1
   length := len(nums) -1 
   for l <= r {
       if nums[l]* nums[l] < nums[r]*nums[r]{
           res[length]= nums[r]*nums[r]
           r--
       } else {
            res[length]= nums[l]* nums[l]
           l++
       }
       length--
   }
   return res
}