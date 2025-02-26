package medium

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Size: 0,
		Head: &Node{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}

	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}

	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	node := &Node{val, curr.Next}
	curr.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
