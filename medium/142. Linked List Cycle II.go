package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle1(head *ListNode) *ListNode {
	visted := make(map[*ListNode]struct{})
	for head != nil {
		if _, found := visted[head]; found {
			return head
		}

		visted[head] = struct{}{}
		head = head.Next
	}

	return nil
}

// Time: O(N)
// Space: O(N)

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return slow
		}
	}

	return nil
}

// Time: 0(N)
// Space: O(1)
