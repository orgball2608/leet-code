package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeLinkedList(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	//find middle pointer
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//reverse list from middle element
	sublist := reverseList(slow)

	//compare two list
	return compare(sublist, head)
}

func compare(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

//Time: 0(N/2 + N + N)
//Space: O(1)
