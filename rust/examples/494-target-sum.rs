use core::num;

fn main() {
    println!("count: {}", Solution::find_target_sum_ways(vec![2,1], 1))
}

struct Solution;

impl Solution {
    /// 给你一个整数数组 nums 和一个整数 target 。
    ///
    ///向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
    ///
    ///例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
    ///
    ///返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
    ///
    #[doc = "https://leetcode-cn.com/problems/target-sum/"]
    pub fn find_target_sum_ways(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() == 1 {
            if nums[0] == target {
                return 1
            } else {
                return 0
            }
        }
        let mut rt = i32::default();
        rt += Self::find_target_sum_ways(nums[1..].to_vec(),   nums[0] + target);
        rt += Self::find_target_sum_ways(nums[1..].to_vec(),   nums[0] + target);
        rt
    }
}
