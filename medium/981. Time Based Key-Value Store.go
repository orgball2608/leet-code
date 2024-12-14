package medium

type TimeMap struct {
	Items map[string][]Item
}

type Item struct {
	Val       string
	TimeStamp int
}

func Constructor() TimeMap {
	return TimeMap{
		Items: make(map[string][]Item, 0),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	_, exist := this.Items[key]
	if !exist {
		this.Items[key] = make([]Item, 0)
	}

	this.Items[key] = append(this.Items[key], Item{
		Val:       value,
		TimeStamp: timestamp,
	})
}

func lowerBound(arr []Item, target int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := (low + high) / 2
		if arr[mid].TimeStamp <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low - 1
}

func (this *TimeMap) Get(key string, timestamp int) string {
	item, exist := this.Items[key]
	if !exist {
		return ""
	}

	if len(item) == 0 {
		return ""
	}

	index := lowerBound(item, timestamp)
	if index < 0 {
		return ""
	}

	return item[index].Val
}

// Time: O(NlogN)
// Space: O(N)
