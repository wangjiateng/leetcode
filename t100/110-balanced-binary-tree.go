package t100

// IsBalanced 平衡二叉树 https://leetcode-cn.com/problems/balanced-binary-tree/
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !IsBalanced(root.Left) || !IsBalanced(root.Right) {
		return false
	}
	c := treeMaxDepth(root.Left) - treeMaxDepth(root.Right)
	return c <= 1 && c >= -1
}

func treeMaxDepth(root *TreeNode) (r int) {
	if root == nil {
		return
	}

	r = treeMaxDepth(root.Left)
	rightdepth := treeMaxDepth(root.Right)

	if rightdepth > r {
		r = rightdepth
	}
	r++

	return
}
