package medium

type SmallestInfiniteSet struct {
	Poped      map[int]struct{}
	MinElement int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		Poped:      make(map[int]struct{}),
		MinElement: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	temp := this.MinElement
	this.MinElement++
	this.Poped[temp] = struct{}{}
	for {
		if _, ok := this.Poped[this.MinElement]; ok {
			this.MinElement++
		} else {
			break
		}
	}

	return temp
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.MinElement {
		this.MinElement = num
	}
	delete(this.Poped, num)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

// Time: O(1)
// Space: O(1)
