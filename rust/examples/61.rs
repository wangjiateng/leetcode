fn main() {
    println!("{}", Solution::is_straight(vec![0, 1, 2, 0, 5]))
}

struct Solution;
impl Solution {
    #[doc = "https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/submissions/"]
    pub fn is_straight(nums: Vec<i32>) -> bool {
        let mut w_count = 0;
        let mut n_nums: Vec<i32> = Vec::default();

        for num in nums {
            if num == 0 {
                w_count += 1;
                continue;
            }
            n_nums.push(num);
        }

        println!("w_count:{} n_nums:{:?}", w_count, n_nums);
        if n_nums.len() <= 1 {
            return true;
        }
        n_nums.sort();
        let mut tmp = n_nums[0];
        for num in n_nums.into_iter().skip(1) {
            let c = num - tmp - 1;
            if c == 0 {
                tmp = num;
                continue;
            }
            if c > 0 && w_count >= c {
                w_count -= c;
                tmp = num;
                continue;
            }
            return false;
        }

        true
    }
}
