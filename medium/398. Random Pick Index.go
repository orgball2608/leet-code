package medium

type Solution struct {
	HashMap map[int][]int
}

func Constructor(nums []int) Solution {
	hashMap := make(map[int][]int)
	for i, num := range nums {
		hashMap[num] = append(hashMap[num], i)
	}
	return Solution{
		HashMap: hashMap,
	}

}

// Time: O(N)
// Space: O(N)

func (this *Solution) Pick(target int) int {
	indexes, _ := this.HashMap[target]
	index := rand.Intn(len(indexes))
	return indexes[index]
}

// Time: O(1)
// Space: O(1)

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
