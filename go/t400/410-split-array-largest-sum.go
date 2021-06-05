package t400

import "fmt"

// SplitArray 分割数组的最大值 https://leetcode-cn.com/problems/split-array-largest-sum/
func SplitArray(nums []int, m int) (r int) {
	return doSplitArray(nums, m, 0, make(map[string]int))
}
func doSplitArray(nums []int, m, s int, t map[string]int) (r int) {

	var has bool
	if r, has = t[fmt.Sprintf("%d-%d", m, s)]; has {
		return
	}

	defer func() {
		t[fmt.Sprintf("%d-%d", m, s)] = r
	}()

	if m == 1 {
		for ; s < len(nums); s++ {
			r += nums[s]
		}
		return
	}

	maxs := []int{}
	var left int
	for i := s + 1; i <= len(nums)-m+1; i++ {
		left += nums[i-1]
		right := doSplitArray(nums, m-1, i, t)
		maxs = append(maxs, max(left, right))
	}
	r = min(maxs)

	return r
}

func min(ay []int) (r int) {

	r = ay[0]
	for i := 1; i < len(ay); i++ {
		if r > ay[i] {
			r = ay[i]
		}
	}

	return
}

func max(x, y int) (r int) {

	r = x
	if x < y {
		r = y
	}
	return
}
