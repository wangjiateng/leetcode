package t650

import "fmt"

// CountBinarySubstrings 计数二进制子串 https://leetcode-cn.com/problems/count-binary-substrings/
func CountBinarySubstrings(s string) (r int) {

	for x := 0; x < len(s)-1; x++ {

		y := x + 1
		ct := 1
		for ; y < len(s); y++ {
			if s[y] != s[y-1] {
				break
			}
			ct++
		}

		ok := true
		for i := 1; i < ct; i++ {
			if y+i >= len(s) || s[y+i] != s[y+i-1] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(string(s[x:y+ct]), ct)
			r += ct
			x = y - 1
		}

	}

	return
}
