package medium

type MinStack struct {
	stack    []int
	minStack []int
	topEl    int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
		topEl:    -1,
	}
}

func (ms *MinStack) Push(val int) {
	if ms.topEl == -1 {
		ms.minStack = append(ms.minStack, val)
	} else {
		ms.minStack = append(ms.minStack, min(ms.minStack[ms.topEl], val))
	}
	ms.stack = append(ms.stack, val)
	ms.topEl++
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:ms.topEl]
	ms.minStack = ms.minStack[:ms.topEl]
	ms.topEl--
	return
}

func (ms *MinStack) Top() int {
	return ms.stack[ms.topEl]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[ms.topEl]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
