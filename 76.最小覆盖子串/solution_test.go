package main

import (
	"math"
	"testing"
)

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



 注意：


 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 如果 s 中存在这样的子串，我们保证它是唯一的答案。




 示例 1：


输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。


 示例 2：


输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。


 示例 3:


输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。



 提示：


 m == s.length
 n == t.length
 1 <= m, n <= 10⁵
 s 和 t 由英文字母组成



进阶：你能设计一个在
o(m+n) 时间内解决此问题的算法吗？

 Related Topics 哈希表 字符串 滑动窗口 👍 2445 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	needMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needMap[t[i]]++
	}

	// 用于快速跟踪needMap中全部为0的情况
	needCount := len(t)

	left := 0
	right := 0
	start := -1
	minSize := math.MaxInt32

	for right < len(s) {
		if needMap[s[right]] > 0 {
			needCount--
		}
		needMap[s[right]]--

		// 当前left, right所组成的窗口已经满足条件，开始缩减窗口左端
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

			// 在破坏窗口成立条件前，看看是否需要更新最优可行性解
			curSize := right - left + 1
			if curSize < minSize {
				minSize = curSize
				start = left
			}

			// 破坏窗口条件，使right继续前行
			needMap[s[left]]++
			needCount++
			left++
		}

		right++
	}

	if start < 0 {
		return ""
	}

	return s[start : start+minSize]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumWindowSubstring(t *testing.T) {
	println(minWindow("ab", "a"))
}
