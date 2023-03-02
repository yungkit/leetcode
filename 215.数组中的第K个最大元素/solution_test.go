package main

import (
	"math/rand"
	"testing"
	"time"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



 示例 1:


输入: [3,2,1,5,6,4], k = 2
输出: 5


 示例 2:


输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4



 提示：


 1 <= k <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴


 Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2073 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSortPartition(nums, 0, len(nums)-1, k)
}

func quickSortPartition(nums []int, left, right, k int) int {

	pivotIndex := partition(nums, left, right)
	if pivotIndex == len(nums)-k {
		return nums[pivotIndex]
	} else if pivotIndex < len(nums)-k {
		return quickSortPartition(nums, pivotIndex+1, right, k)
	}
	return quickSortPartition(nums, left, pivotIndex-1, k)
}

func partition(nums []int, left, right int) int {
	randomIndex := left + rand.Intn(right-left+1)
	nums[left], nums[randomIndex] = nums[randomIndex], nums[left]

	i := left + 1
	j := right

	pivot := nums[left]

	for {
		for i <= right && nums[i] < pivot {
			i++
		}

		for j > left && nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	nums[left], nums[j] = nums[j], nums[left]
	return j
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthLargestElementInAnArray(t *testing.T) {

}
