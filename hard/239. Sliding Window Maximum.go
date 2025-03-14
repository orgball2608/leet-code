package hard

type MaxQueue struct {
	Data []int
}

func NewMaxQueue() MaxQueue {
	return MaxQueue{
		Data: []int{},
	}
}

func (m *MaxQueue) EnQueue(x int) {
	for len(m.Data) > 0 && m.Data[len(m.Data)-1] < x {
		m.Data = m.Data[:len(m.Data)-1]
	}
	m.Data = append(m.Data, x)
}

func (m *MaxQueue) DeQueue(x int) int {
	if len(m.Data) == 0 {
		return -1
	}
	first := m.Data[0]
	if x == first {
		m.Data = m.Data[1:]
	}
	return first
}

func (m *MaxQueue) GetMax() int {
	if len(m.Data) > 0 {
		return m.Data[0]
	}
	return -1
}

func maxSlidingWindow(nums []int, k int) []int {
	maxQueue := NewMaxQueue()
	for i := 0; i < k; i++ {
		maxQueue.EnQueue(nums[i])
	}

	n := len(nums)
	result := make([]int, n-k+1)
	for i := k; i < n; i++ {
		result[i-k] = maxQueue.GetMax()
		maxQueue.EnQueue(nums[i])
		maxQueue.DeQueue(nums[i-k])
	}
	result[n-k] = maxQueue.GetMax()
	return result
}

// Time: O(N)
// Space: O(N)
