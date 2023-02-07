package main

import (
	"strings"
	"testing"
)

/**
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。

 本题中，将空字符串定义为有效的 回文串 。



 示例 1:


输入: s = "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串

 示例 2:


输入: s = "race a car"
输出: false
解释："raceacar" 不是回文串



 提示：


 1 <= s.length <= 2 * 10⁵
 字符串 s 由 ASCII 字符组成





 注意：本题与主站 125 题相同： https://leetcode-cn.com/problems/valid-palindrome/

 Related Topics 双指针 字符串 👍 42 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {

	n := len(s)
	if n == 0 {
		return true
	}

	pureAlNumStrBytes := make([]byte, 0)
	for i := 0; i < n; i++ {
		if isAlNum(s[i]) {
			pureAlNumStrBytes = append(pureAlNumStrBytes, s[i])
		}
	}

	pureAlNumStr := string(pureAlNumStrBytes)
	pureAlNumStr = strings.ToLower(pureAlNumStr)
	n = len(pureAlNumStr)
	for i := 0; i < n; i++ {
		if pureAlNumStr[i] != pureAlNumStr[n-i-1] {
			return false
		}
	}

	return true
}

func isAlNum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

//leetcode submit region end(Prohibit modification and deletion)

func TestXltzEq(t *testing.T) {

}
