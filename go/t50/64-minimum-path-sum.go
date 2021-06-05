package t50

import "fmt"

// MinPathSum 最小路径和 https://leetcode-cn.com/problems/minimum-path-sum/
func MinPathSum(grid [][]int) (r int) {
	return minPath(grid, 0, 0, map[string]int{})
}

func minPath(grid [][]int, x, y int, t map[string]int) (r int) {

	tk := fmt.Sprintf("%d-%d", x, y)
	var has bool
	if r, has = t[tk]; has {
		return r
	}

	defer func() {
		t[tk] = r
	}()

	if x == len(grid)-1 && y == len(grid[x])-1 {
		return grid[x][y]
	}

	if x == len(grid)-1 {
		return grid[x][y] + minPath(grid, x, y+1, t)
	}

	if y == len(grid[x])-1 {
		return grid[x][y] + minPath(grid, x+1, y, t)
	}

	min := minPath(grid, x, y+1, t)
	d := minPath(grid, x+1, y, t)

	if min > d {
		min = d
	}

	return grid[x][y] + min
}
