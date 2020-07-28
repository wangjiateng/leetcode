package t100

// MaxDepth 二叉树的最大深度 https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func MaxDepth(root *TreeNode) (r int) {

	if root == nil {
		return 0
	}

	r = MaxDepth(root.Left)
	rd := MaxDepth(root.Right)
	if r < rd {
		r = rd
	}

	r++
	return
}

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
