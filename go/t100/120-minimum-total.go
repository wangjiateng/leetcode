package t100

import (
	"math"
	"strconv"
)

// MinimumTotal 三角形最小路径和 https://leetcode-cn.com/classic/problems/triangle/description/
func MinimumTotal(triangle [][]int) (r int) {
	r = getMinTotal(0, 0, triangle, map[string]int{})
	return
}

func getMinTotal(ci, li int, triangle [][]int, tm map[string]int) (r int) {

	if tr, has := tm[strconv.Itoa(ci)+"_"+strconv.Itoa(li)]; has {
		r = tr
		return
	}

	r = triangle[ci][li]

	if ci == len(triangle)-1 { // 最后一层
		return
	}

	lt, rt := getMinTotal(ci+1, li, triangle, tm), getMinTotal(ci+1, li+1, triangle, tm)
	r += int(math.Min(float64(lt), float64(rt)))

	tm[strconv.Itoa(ci)+"_"+strconv.Itoa(li)] = r

	return
}
