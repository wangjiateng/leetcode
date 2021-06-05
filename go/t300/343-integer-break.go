package t300

// IntegerBreak 整数拆分 https://leetcode-cn.com/problems/integer-break/
func IntegerBreak(n int) (r int) {
	return doIntegerBreak(n, false, map[int]int{})
}

func doIntegerBreak(n int, flag bool, mp map[int]int) (r int) {

	if n <= 2 {
		r = 1
		if flag {
			r = 2
		}
		return
	}

	var has bool
	if r, has = mp[n]; has {
		return
	}

	defer func() {
		mp[n] = r
	}()

	if flag {
		r = n
	}
	for i := 1; i < n/2+1; i++ {
		max := i * doIntegerBreak(n-i, true, mp)
		if r < max {
			r = max
		}
	}

	return
}
