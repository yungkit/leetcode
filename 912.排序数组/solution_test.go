package main

import (
	"math/rand"
	"testing"
	"time"
)

/**
给你一个整数数组 nums，请你将该数组升序排列。






 示例 1：


输入：nums = [5,2,3,1]
输出：[1,2,3,5]


 示例 2：


输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]




 提示：


 1 <= nums.length <= 5 * 10⁴
 -5 * 10⁴ <= nums[i] <= 5 * 10⁴


 Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 787 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {

	length := len(nums)
	if length <= 1 {
		return nums
	}

	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, length-1)

	return nums

}

func quickSort(nums []int, left, right int) {

	if left >= right {
		return
	}
	pivotIndex := partition(nums, left, right)
	quickSort(nums, left, pivotIndex-1)
	quickSort(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	// 采用随机取pivot的方式
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

func TestSortAnArray(t *testing.T) {
	sortArray([]int{5, 2, 3, 1})
}
