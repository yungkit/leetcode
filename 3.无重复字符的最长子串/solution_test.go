package main

import (
	"testing"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



 示例 1:


输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


 示例 2:


输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


 示例 3:


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。




 提示：


 0 <= s.length <= 5 * 10⁴
 s 由英文字母、数字、符号和空格组成


 Related Topics 哈希表 字符串 滑动窗口 👍 8661 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	ans := 1
	right := -1

	// 用于left-right所表示的范围内判断是否存在重复字符
	m := make(map[byte]int)
	strLen := len(s)
	for left := 0; left < strLen; left++ {
		if left != 0 {
			delete(m, s[left-1])
		}

		// 要考虑第一步的情况，就知道为什么right初始化为-1，且右移过程中，下标为right+1了
		for right+1 < strLen && m[s[right+1]] == 0 {
			m[s[right+1]] = 1
			right++
		}

		windowSize := right - left + 1
		if windowSize > ans {
			ans = windowSize
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	println(lengthOfLongestSubstring("abcabcbb"))
}
