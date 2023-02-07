package main

import (
	"testing"
)

/**
给你一个由小写英文字母组成的回文字符串 palindrome ，请你将其中 一个 字符用任意小写英文字母替换，使得结果字符串的 字典序最小 ，且 不是 回文串。


 请你返回结果字符串。如果无法做到，则返回一个 空串 。

 如果两个字符串长度相同，那么字符串 a 字典序比字符串 b 小可以这样定义：在 a 和 b 出现不同的第一个位置上，字符串 a 中的字符严格小于 b 中的对应
字符。例如，"abcc” 字典序比 "abcd" 小，因为不同的第一个位置是在第四个字符，显然 'c' 比 'd' 小。

 示例 1：


输入：palindrome = "abccba"
输出："aaccba"
解释：存在多种方法可以使 "abccba" 不是回文，例如 "zbccba", "aaccba", 和 "abacba" 。
在所有方法中，"aaccba" 的字典序最小。

 示例 2：


输入：palindrome = "a"
输出：""
解释：不存在替换一个字符使 "a" 变成非回文的方法，所以返回空字符串。



 提示：


 1 <= palindrome.length <= 1000
 palindrome 只包含小写英文字母。


 Related Topics 贪心 字符串 👍 49 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func breakPalindrome(palindrome string) string {
	if len(palindrome) <= 1 {
		return ""
	}

	bytes := []byte(palindrome)
	n := len(bytes)

	/**
	  首先如果字符串长度为奇数，则字符串中间的那个字符无论怎么改，字符串都是回文串。 如：aba，b字符无论怎么改，字符串都还是回文串。

	  回文串前半段和后半段是相互对应的，因此只要遍历一半就好了。

	  首先遍历前半段，遇到不为a的字符就直接将其替换成a，然后直接return结果。 如果前半段都是a，则说明后半段也都是a，说明字符串要么类似aabaa，要么类似aaaaaa。 直接将最后1个字符改成b就好了。
	*/

	// 注意结束条件，找几个3，4，5的case试试就好了，对于奇数最中间那个，怎么替换都是没用的，所以(n-1)/2是不需要包括的
	// 直接跳到最后的处理
	for i := 0; i <= (n-2)/2; i++ {

		// 找到第一个不为a的替换为a
		if bytes[i] > 'a' {
			bytes[i] = 'a'
			return string(bytes)
		}
	}

	// 说明除了中间外，都是a, 直接使最后一位变为b
	bytes[n-1] = 'b'

	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBreakAPalindrome(t *testing.T) {
	println(breakPalindrome("aba"))
}
