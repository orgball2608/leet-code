package medium

type MyCircularQueue struct {
	Data       []int
	FrontIndex int
	RearIndex  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Data:       make([]int, k),
		FrontIndex: -1,
		RearIndex:  -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.FrontIndex = 0
	}

	this.RearIndex = (this.RearIndex + 1) % len(this.Data)
	this.Data[this.RearIndex] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.FrontIndex == this.RearIndex {
		this.FrontIndex = -1
		this.RearIndex = -1
	} else {
		this.FrontIndex = (this.FrontIndex + 1) % len(this.Data)
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Data[this.FrontIndex]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Data[this.RearIndex]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.FrontIndex == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.RearIndex+1)%len(this.Data) == this.FrontIndex
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
