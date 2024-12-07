package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	checkLeft := isSameTree(p.Left, q.Left)
	checkRight := isSameTree(p.Right, q.Right)

	return checkLeft && checkRight
}

// Time: O(N)
// Space: O(N)
