package medium

type RandomizedSet struct {
	Data     []int
	IndexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Data:     []int{},
		IndexMap: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.IndexMap[val]
	if ok {
		return false
	}

	this.Data = append(this.Data, val)
	this.IndexMap[val] = len(this.Data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.IndexMap[val]
	if !ok {
		return false
	}

	last := len(this.Data) - 1
	if index != last {
		this.Data[index] = this.Data[last]
		this.IndexMap[this.Data[index]] = index
	}
	this.Data = this.Data[:len(this.Data)-1]
	delete(this.IndexMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.Data[rand.Intn(len(this.Data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
