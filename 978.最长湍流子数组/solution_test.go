package main

import (
	"testing"
)

/**
给定一个整数数组 arr ，返回 arr 的 最大湍流子数组的长度 。

 如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是 湍流子数组 。

 更正式地来说，当 arr 的子数组 A[i], A[i+1], ..., A[j] 满足仅满足下列条件时，我们称其为湍流子数组：


 若 i <= k < j ：



 当 k 为奇数时， A[k] > A[k+1]，且
 当 k 为偶数时，A[k] < A[k+1]；


 或 若 i <= k < j ：

 当 k 为偶数时，A[k] > A[k+1] ，且
 当 k 为奇数时， A[k] < A[k+1]。






 示例 1：


输入：arr = [9,4,2,10,7,8,8,1,9]
输出：5
解释：arr[1] > arr[2] < arr[3] > arr[4] < arr[5]

 示例 2：


输入：arr = [4,8,12,16]
输出：2


 示例 3：


输入：arr = [100]
输出：1




 提示：


 1 <= arr.length <= 4 * 10⁴
 0 <= arr[i] <= 10⁹


 Related Topics 数组 动态规划 滑动窗口 👍 223 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxTurbulenceSize(arr []int) int {

	n := len(arr)
	if n <= 1 {
		return n
	}

	left := 0
	right := 0
	ans := 0
	for right < n-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestTurbulentSubarray(t *testing.T) {

}
