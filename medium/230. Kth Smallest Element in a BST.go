package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var result int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		k--
		if k == 0 {
			result = root.Val
			return
		}
		dfs(root.Right)
	}

	dfs(root)

	return result
}

// Time: O(N)
// Space: O(H)
