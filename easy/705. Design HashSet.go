package easy

type MyHashSet struct {
	Data [][]int
}

const (
	_capacity = 10000
)

func Constructor() MyHashSet {
	return MyHashSet{
		Data: make([][]int, _capacity),
	}
}

func (this *MyHashSet) Add(key int) {
	index := getIndex(key)
	bucketIndex := exist(key, this.Data[index])
	if bucketIndex == -1 {
		this.Data[index] = append(this.Data[index], key)
	}
}

// Time: O(N/K)

func (this *MyHashSet) Remove(key int) {
	index := getIndex(key)
	bucketIndex := exist(key, this.Data[index])
	if bucketIndex != -1 {
		this.Data[index] = remove(this.Data[index], bucketIndex)
	}
}

// Time: O(N/K)

func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func getIndex(key int) int {
	return key % _capacity
}

func exist(key int, data []int) int {
	for i, val := range data {
		if key == val {
			return i
		}
	}

	return -1
}

// Time: O(N/K)

func (this *MyHashSet) Contains(key int) bool {
	index := getIndex(key)
	if exist(key, this.Data[index]) != -1 {
		return true
	}

	return false
}

// Time: O(N/K)
// Space: O(N)

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
