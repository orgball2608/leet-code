package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode, low, high int64) bool
	dfs = func(root *TreeNode, low, high int64) bool {
		if root == nil {
			return true
		}

		if int64(root.Val) >= high || int64(root.Val) <= low {
			return false
		}

		return dfs(root.Left, low, int64(root.Val)) && dfs(root.Right, int64(root.Val), high)
	}

	return dfs(root, math.MinInt64, math.MaxInt64)
}

// Time: O(N)
// Space: O(H)
