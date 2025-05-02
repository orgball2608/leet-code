package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return similarTree(root.Left, root.Right)
}

func similarTree(u, v *TreeNode) bool {
	if u == nil {
		return v == nil
	}

	if v == nil {
		return u == nil
	}

	if v.Val != u.Val {
		return false
	}

	return similarTree(u.Left, v.Right) && similarTree(v.Left, u.Right)
}

// Time: O(n)
// Space: O(H)
