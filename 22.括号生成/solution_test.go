package main

import (
	"fmt"
	"testing"
)

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



 示例 1：


输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]


 示例 2：


输入：n = 1
输出：["()"]




 提示：


 1 <= n <= 8


 Related Topics 字符串 动态规划 回溯 👍 3045 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

var ans []string

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	ans = make([]string, 0)

	getParnthesis("", n, n)

	return ans
}

func getParnthesis(str string, left, right int) {
	if left == 0 && right == 0 {
		ans = append(ans, str)
		return
	}

	// 剩余左右括号数相等，下一个只能用左括号
	if left == right {
		getParnthesis(str+"(", left-1, right)
	} else {
		// 这里必然是left < right
		// 既可以用左括号，也可以用右括号
		if left > 0 {
			getParnthesis(str+"(", left-1, right)
		}
		getParnthesis(str+")", left, right-1)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGenerateParentheses(t *testing.T) {
	fmt.Printf("%+v \n", generateParenthesis(1))
}
