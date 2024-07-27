package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//dummy node
	dummyNode := &ListNode{0, head}
	fast, slow := dummyNode, dummyNode
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
}

//Time: O(N)
//Space: O(1)
