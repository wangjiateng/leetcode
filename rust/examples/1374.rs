fn main() {
    println!("{}", Solution::generate_the_string(7))
}

struct Solution;
impl Solution {
    #[doc = "https://leetcode-cn.com/problems/generate-a-string-with-characters-that-have-odd-counts/"]
    pub fn generate_the_string(n: i32) -> String {
        let mut rt_vec: Vec<u8> = Vec::default();
        match n {
            1 => return 'a'.to_string(),
            2 => return "ab".to_string(),
            _ => {}
        };
        let mut h = n | 1;
        if h > n {
            h = h - 2;
        }
        for _ in 0..h {
            rt_vec.push(b'a');
        }
        if n % h == 1 {
            rt_vec.push(b'b');
        }
        return String::from_utf8(rt_vec).unwrap();
    }
}
