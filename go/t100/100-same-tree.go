package t100

// IsSameTree 相同的树 https://leetcode-cn.com/problems/same-tree/
func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if IsSameTree(p.Right, q.Right) && IsSameTree(p.Left, q.Left) {
		return true
	}

	return false
}
