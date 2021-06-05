package t350

// IsSubsequence 判断子序列 https://leetcode-cn.com/problems/is-subsequence/
func IsSubsequence(s string, t string) (r bool) {

	ti := 0
	for _, s := range s {
		var ok bool
		for ; ti < len(t); ti++ {
			ts := []rune(t)[ti]

			if ts == s {
				ti++
				ok = true
				break
			}
		}

		if !ok {
			return
		}
	}

	r = true

	return
}
