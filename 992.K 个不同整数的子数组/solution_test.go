package main

import (
	"testing"
)

/**
给定一个正整数数组 nums和一个整数 k ，返回 num 中 「好子数组」 的数目。

 如果 nums 的某个子数组中不同整数的个数恰好为 k，则称 nums 的这个连续、不一定不同的子数组为 「好子数组 」。


 例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。


 子数组 是数组的 连续 部分。



 示例 1：


输入：nums = [1,2,1,2,3], k = 2
输出：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].


 示例 2：


输入：nums = [1,2,1,3,4], k = 3
输出：3
解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].




 提示：


 1 <= nums.length <= 2 * 10⁴
 1 <= nums[i], k <= nums.length


 Related Topics 数组 哈希表 计数 滑动窗口 👍 431 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

// https://leetcode.cn/problems/subarrays-with-k-different-integers/solutions/598241/cong-zui-jian-dan-de-wen-ti-yi-bu-bu-tuo-7f4v/?orderBy=most_votes
func subarraysWithKDistinct(nums []int, k int) int {

	if len(nums) < k || k == 0 {
		return 0
	}

	return atMostKDistinct(nums, k) - atMostKDistinct(nums, k-1)

}

func atMostKDistinct(nums []int, k int) int {
	ans := 0
	countMap := map[int]int{}
	n := len(nums)

	left := 0
	right := 0
	for right < n {
		countMap[nums[right]]++

		for len(countMap) > k {
			countMap[nums[left]]--
			if countMap[nums[left]] == 0 {
				delete(countMap, nums[left])
			}
			left++
		}
		ans += right - left + 1
		right++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubarraysWithKDifferentIntegers(t *testing.T) {

}
