package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	set := make(map[int]struct{})
	dummy := &ListNode{Next: head}
	prev, curr := dummy, head
	for curr != nil {
		if _, found := set[curr.Val]; found {
			prev.Next = curr.Next
		} else {
			set[curr.Val] = struct{}{}
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}

// Time: O(N)
// Spaace: O(N)

func deleteDuplicates2(head *ListNode) *ListNode {
	temp := head
	for temp != nil && temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head
}

// Time: O(N)
// Space: O(1)
