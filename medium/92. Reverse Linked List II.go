package medium

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{0, head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	leftNode := prev.Next
	rightNode := leftNode
	for i := 0; i < right-left; i++ {
		rightNode = rightNode.Next
	}

	after := rightNode.Next
	rightNode.Next = nil
	newLeft, newRight := reverse(leftNode)

	prev.Next = newLeft
	newRight.Next = after

	return dummy.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev, head
}

// Time: O(N)
// Space: O(1)
