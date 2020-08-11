package t100

import (
	"fmt"
	"strconv"
	"strings"
)

// Solve 被围绕的区域 https://leetcode-cn.com/problems/surrounded-regions/
func Solve(board [][]byte) {

	omp := map[string]struct{}{}
	for x := range board {
		for y, v := range board[x] {
			if v == 'O' {
				omp[fmt.Sprintf("%d_%d", x, y)] = struct{}{}
			}
		}
	}

	bay := map[[2]int]struct{}{}
	for x := range board {
		for y, v := range board[x] {
			if v == 'O' && ((x == 0 || x == len(board)-1) || (y == 0 || y == len(board[x])-1)) {
				bay[[2]int{x, y}] = struct{}{}
			}
		}
	}

	for p := range bay {
		cleanomap(p[0], p[1], omp)
	}

	for pstr := range omp {
		ss := strings.Split(pstr, "_")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])
		board[x][y] = 'X'
	}

}

func cleanomap(x, y int, omp map[string]struct{}) {
	k := fmt.Sprintf("%d_%d", x, y)
	if _, has := omp[fmt.Sprintf("%d_%d", x, y)]; has {
		delete(omp, k)
		friends := [][]int{
			{x - 1, y},
			{x, y - 1}, {x, y + 1},
			{x + 1, y},
		}
		for _, p := range friends {
			cleanomap(p[0], p[1], omp)
		}
	}
}
