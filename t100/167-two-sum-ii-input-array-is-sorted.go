package t100

//TwoSum 两数之和 II - 输入有序数组 https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSum(numbers []int, target int) (r []int) {

	si, ei := 0, len(numbers)-1
	sum := numbers[si] + numbers[ei]
	for sum != target {
		if sum > target {
			ei--
		} else if sum < target {
			si++
		}
		sum = numbers[si] + numbers[ei]
	}

	return []int{si + 1, ei + 1}
}
