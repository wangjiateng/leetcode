package m800

// FindMagicIndex 面试题 08.03. 魔术索引 https://leetcode-cn.com/problems/magic-index-lcci/
func FindMagicIndex(nums []int) int {

	for i, num := range nums {
		if num == i {
			return i
		}
	}

	return -1
}
