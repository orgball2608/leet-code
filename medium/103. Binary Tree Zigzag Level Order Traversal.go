package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
		leftToRight = !leftToRight
	}

	return result
}

// Time: O(N)
// Space: O(N)
