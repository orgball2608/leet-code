package easy

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	prev := 0
	hasPrev := false
	minDiff := math.MaxInt32

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if hasPrev {
			diff := abs(node.Val - prev)
			minDiff = min(minDiff, diff)
		}
		prev = node.Val
		hasPrev = true

		inorder(node.Right)
	}

	inorder(root)
	return minDiff
}

// Time: O(N) where N is the number of nodes in the tree
// Space: O(H) where H is the height of the tree
