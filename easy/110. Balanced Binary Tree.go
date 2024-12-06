package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, balanced := checkBalance(root)
	return balanced
}

func checkBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	left, leftIsBalance := checkBalance(root.Left)
	right, rightIsBalance := checkBalance(root.Right)
	if !leftIsBalance || !rightIsBalance || abs(left-right) > 1 {
		return 0, false
	}

	return 1 + Max(left, right), true
}

// Time: O(N)
// Space: O(N)
