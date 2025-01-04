package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := getLeaf(root1)
	leaf2 := getLeaf(root2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true
}

func getLeaf(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
	} else {
		result = append(result, getLeaf(root.Left)...)
		result = append(result, getLeaf(root.Right)...)
	}

	return result
}

// Time: O(N1+N2+K) K is leaf count
// Space: O(H1+H2+K)
