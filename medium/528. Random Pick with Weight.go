package medium

type Solution struct {
	Weight []int
	Sum    int
}

func Constructor(w []int) Solution {
	sum := 0
	for _, val := range w {
		sum += val
	}

	return Solution{
		Weight: w,
		Sum:    sum,
	}
}

func (this *Solution) PickIndex() int {
	randInt := rand.Intn(this.Sum)
	sum := 0
	for i, val := range this.Weight {
		sum += val
		if sum > randInt {
			return i
		}
	}
	return len(this.Weight) - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

// Solution 1:
// Time: O(N)
// Space: O(N)

type Solution struct {
	PrefixSum []int
	Sum       int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, 0, len(w))
	sum := 0
	for _, val := range w {
		sum += val
		prefixSum = append(prefixSum, sum)
	}

	return Solution{
		PrefixSum: prefixSum,
		Sum:       sum,
	}
}

func (this *Solution) PickIndex() int {
	randInt := rand.Intn(this.Sum)
	index := sort.Search(len(this.PrefixSum), func(i int) bool {
		return this.PrefixSum[i] > randInt
	})

	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

// Solution 2:
// Time: O(N) for init and O(logN) for pickIndex
// Space: O(N)
