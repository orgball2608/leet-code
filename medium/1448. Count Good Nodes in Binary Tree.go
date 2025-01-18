package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}

		good := 0
		if node.Val >= maxVal {
			good = 1
		}

		maxVal = max(maxVal, node.Val)

		good += dfs(node.Left, maxVal)
		good += dfs(node.Right, maxVal)

		return good
	}

	return dfs(root, root.Val)
}

// Time: 0(N)
// Space: O(H)
