package t50

// GenerateTrees 不同的二叉搜索树 II https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
func GenerateTrees(n int) (r []*TreeNode) {
	if n < 1 {
		return []*TreeNode{}
	}
	return makeTree(1, n)
}

func makeTree(start, end int) []*TreeNode {

	if start > end {
		return []*TreeNode{nil}
	}

	var r []*TreeNode
	for i := start; i <= end; i++ {
		for _, left := range makeTree(start, i-1) {
			for _, right := range makeTree(i+1, end) {
				r = append(r, &TreeNode{i, left, right})
			}
		}
	}

	return r
}

//TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
