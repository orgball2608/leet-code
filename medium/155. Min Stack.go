package medium

type MinStack struct {
	Stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{
		Stack:    []int{},
		MinStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 || val <= this.MinStack[len(this.MinStack)-1] {
		this.MinStack = append(this.MinStack, val)
	}
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	if this.Stack[len(this.Stack)-1] == this.MinStack[len(this.MinStack)-1] {
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.MinStack) == 0 {
		return 0
	}
	return this.MinStack[len(this.MinStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
