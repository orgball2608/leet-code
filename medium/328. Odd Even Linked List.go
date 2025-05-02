package medium

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	curr := head
	var prev *ListNode = dummy

	newHead := &ListNode{}
	oddCur := newHead

	count := 0
	for curr != nil {
		next := curr.Next
		if count%2 == 1 {
			prev.Next = curr.Next

			curr.Next = nil
			oddCur.Next = curr
			oddCur = oddCur.Next
		} else {
			prev = curr
		}
		curr = next
		count++
	}

	prev.Next = newHead.Next

	return dummy.Next
}

// Time Complexity: O(n)
// Space Complexity: O(1)
