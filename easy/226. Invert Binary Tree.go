package easy

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	left := root.Left
	right := root.Right
	root.Left = right
	root.Right = left

	traverse(root.Left)
	traverse(root.Right)
}

//Time: 0(N)
//Space: 0(N)
