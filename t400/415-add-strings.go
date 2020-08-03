package t400

import (
	"fmt"
	"strconv"
)

// AddStrings 字符串相加 https://leetcode-cn.com/problems/add-strings/
func AddStrings(num1 string, num2 string) (r string) {

	maxLen := len(num1)
	if len(num2) > maxLen {
		maxLen = len(num2)
	}

	var j int
	for i := 0; i < maxLen; i++ {
		n1, n2 := getInt(num1, i), getInt(num2, i)
		v := n1 + n2 + j
		if v >= 10 {
			j = 1
		} else {
			j = 0
		}
		v = v % 10
		r = fmt.Sprintf("%d%s", v, r)
	}

	if j == 1 {
		r = fmt.Sprintf("1%s", r)
	}

	return
}

func getInt(num string, index int) (r int) {

	i := len(num) - 1 - index
	if i < 0 {
		r = 0
	} else {
		r, _ = strconv.Atoi(num[i : i+1])
	}

	return
}
