package medium

type LRUNode struct {
	Key, Val   int
	Next, Prev *LRUNode
}

type LRUCache struct {
	hash       map[int]*LRUNode
	head, tail *LRUNode
	cap        int
	len        int
}

func ConstructorLRU(capacity int) LRUCache {
	head, tail := &LRUNode{}, &LRUNode{}

	head.Next = tail
	tail.Prev = head
	return LRUCache{
		hash: make(map[int]*LRUNode, capacity),
		head: head,
		tail: tail,
		cap:  capacity,
		len:  0,
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.hash[key]
	if !ok {
		return -1
	}

	c.MoveNodeToTop(v)

	return v.Val
}

func (c *LRUCache) MoveNodeToTop(node *LRUNode) {
	if node == c.head.Next {
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	first := c.head.Next
	node.Next = first
	node.Prev = c.head
	c.head.Next = node
	first.Prev = node
}

func (c *LRUCache) AddNodeToTop(node *LRUNode) {
	node.Next = c.head.Next
	node.Prev = c.head
	c.head.Next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) Put(key int, value int) {
	if c.Get(key) != -1 {
		c.head.Next.Val = value
		return
	}

	if c.len >= c.cap {
		delete(c.hash, c.tail.Prev.Key)
		c.tail.Prev = c.tail.Prev.Prev
		c.tail.Prev.Next = c.tail
		c.len--
	}

	newNode := &LRUNode{Key: key, Val: value}
	c.AddNodeToTop(newNode)
	c.hash[key] = newNode
	c.len++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
