package t350

// Intersect 两个数组的交集 II https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
func Intersect(nums1 []int, nums2 []int) (r []int) {

	mp := make(map[int]int)
	for _, num := range nums1 {
		if ct, has := mp[num]; has {
			mp[num] = ct + 1
		} else {
			mp[num] = 1
		}
	}

	for _, num := range nums2 {
		ct, has := mp[num]
		if !has || ct == 0 {
			continue
		}

		mp[num] = ct - 1
		r = append(r, num)

	}

	return
}
