package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var results [][]int
	if root == nil {
		return results
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lenLevel := len(queue)
		var values []int
		for ; lenLevel > 0; lenLevel-- {
			node := queue[0]
			queue = queue[1:]

			values = append(values, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if len(values) > 0 {
			results = append(results, values)
		}
	}

	return results
}
