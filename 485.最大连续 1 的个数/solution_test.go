package main

import (
	"testing"
)

/**
给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。



 示例 1：


输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.


 示例 2:


输入：nums = [1,0,1,1,0,1]
输出：2




 提示：


 1 <= nums.length <= 10⁵
 nums[i] 不是 0 就是 1.


 Related Topics 数组 👍 377 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

// 解法1
func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	maxCnt := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			maxCnt = max(maxCnt, cnt)
			cnt = 0
		}
	}

	maxCnt = max(maxCnt, cnt)
	return maxCnt
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 解法2
func findMaxConsecutiveOnes2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	left := 0
	right := 0
	ans := 0
	for right < n {

		if nums[right] == 0 {
			right++
			continue
		}

		for left <= right {
			curSize := right - left + 1
			if curSize == sum(nums[left:right+1]) {
				if curSize > ans {
					ans = curSize
				}
				break
			}
			left++
		}

		right++
	}

	return ans

}

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaxConsecutiveOnes(t *testing.T) {

}
