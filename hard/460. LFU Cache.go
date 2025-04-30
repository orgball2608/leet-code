package hard

type Node struct {
	Key  int
	Val  int
	Freq int
	Prev *Node
	Next *Node
}

type DoubleList struct {
	Head *Node
	Tail *Node
}

func NewDoublist() *DoubleList {
	tail := &Node{}
	head := &Node{}
	head.Next = tail
	tail.Prev = head
	return &DoubleList{
		Head: head,
		Tail: tail,
	}
}

func (dll *DoubleList) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *DoubleList) addToFront(node *Node) {
	temp := l.Head.Next
	l.Head.Next = node
	node.Next = temp
	node.Prev = l.Head
	temp.Prev = node
}

func (l *DoubleList) removeLast() *Node {
	if l.Head.Next == l.Tail {
		return nil
	}

	node := l.Tail.Prev
	prev := node.Prev
	prev.Next = l.Tail
	l.Tail.Prev = prev
	return node
}

func (l *DoubleList) isEmpty() bool {
	if l.Head.Next == l.Tail {
		return true
	}

	return false
}

type LFUCache struct {
	Cap     int
	Len     int
	MinFreq int
	Data    map[int]*Node
	Freqs   map[int]*DoubleList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Cap:     capacity,
		Len:     0,
		MinFreq: 0,
		Data:    make(map[int]*Node, capacity),
		Freqs:   make(map[int]*DoubleList, capacity),
	}
}

func (this *LFUCache) updateFreq(node *Node) {
	list := this.Freqs[node.Freq]
	list.remove(node)
	if list.isEmpty() && node.Freq == this.MinFreq {
		this.MinFreq++
	}
	node.Freq++

	freqList, ok := this.Freqs[node.Freq]
	if !ok {
		freqList = NewDoublist()
		this.Freqs[node.Freq] = freqList
	}
	freqList.addToFront(node)

}

func (this *LFUCache) Get(key int) int {
	v, ok := this.Data[key]
	if !ok {
		return -1
	}

	this.updateFreq(v)
	return v.Val
}

// Time: O(1)
// Space: O(n)

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.Data[key]
	if ok {
		node.Val = value
		this.updateFreq(node)
		return
	}

	if this.Len == this.Cap {
		removeList := this.Freqs[this.MinFreq]
		removeNode := removeList.removeLast()
		delete(this.Data, removeNode.Key)
		this.Len--
	}

	newNode := &Node{key, value, 1, nil, nil}
	if this.Freqs[1] == nil {
		this.Freqs[1] = NewDoublist()
	}
	this.Freqs[1].addToFront(newNode)
	this.Data[key] = newNode
	this.MinFreq = 1
	this.Len++
}

// Time: O(1)
// Space: O(n)

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
