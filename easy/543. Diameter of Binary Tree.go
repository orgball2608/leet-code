package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var maxLen = 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := dfs(node.Left), dfs(node.Right)
		maxLen = max(maxLen, left+right)
		return 1 + max(left, right)
	}

	_ = dfs(root)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Time Complexity: O(n)
// Space Complexity: O(n)
