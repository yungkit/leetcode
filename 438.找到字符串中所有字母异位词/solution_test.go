package main

import (
	"testing"
)

/**
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



 示例 1:


输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。


 示例 2:


输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。




 提示:


 1 <= s.length, p.length <= 3 * 10⁴
 s 和 p 仅包含小写字母


 Related Topics 哈希表 字符串 滑动窗口 👍 1142 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {

	if len(s) < len(p) {
		return nil
	}

	needMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		needMap[p[i]]++
	}

	needCount := len(p)

	left := 0
	right := 0

	var ans []int
	for right < len(s) {
		c := s[right]

		// 推进right
		if needMap[c] > 0 {
			needCount--
		}

		// 对应字符的count--
		needMap[c]--

		// 当前窗口开始满足题目限定条件
		if needCount == 0 {
			for {
				// 这里表示left再继续右移就要破坏窗口成立条件了
				if needMap[s[left]] == 0 {
					break
				}

				// 释放就是加
				needMap[s[left]]++
				left++
			}

			// 在满足各个字符串都对了情况下，再剔除掉一些多余的字符
			if (right - left + 1) == len(p) {
				ans = append(ans, left)
			}

			needMap[s[left]]++
			needCount++
			left++
		}

		right++
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllAnagramsInAString(t *testing.T) {

}
