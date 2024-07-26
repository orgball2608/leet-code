package medium

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	middle := findMiddle(head)
	secondList := middle.Next
	middle.Next = nil
	reversedSecondList := reverse(secondList)
	a := head
	b := reversedSecondList
	for b != nil {
		aNext := a.Next
		bNext := b.Next
		a.Next = b
		b.Next = aNext
		a = aNext
		b = bNext
	}
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var revHead *ListNode

	for head != nil {
		temp := head.Next
		head.Next = revHead
		revHead = head
		head = temp
	}

	return revHead
}

//Time: O(N)
//Space: O(1)
