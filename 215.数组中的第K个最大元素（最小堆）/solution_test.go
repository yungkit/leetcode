package main

import (
	"testing"
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

	if len(nums) < k {
		return -1
	}

	// 维护一个大小只有k的最小堆
	minHeap := buildMinHeap(k)
	for i := range nums {
		minHeap.Push(nums[i])
	}

	return minHeap.Pop()
}

type MinKthHeap struct {
	nums     []int
	size     int
	capacity int
}

func buildMinHeap(capacity int) *MinKthHeap {
	nums := make([]int, capacity)
	return &MinKthHeap{nums: nums, capacity: capacity}
}

func (h *MinKthHeap) Push(val int) {

	if h.size >= h.capacity {
		// 对比即将要插入的数据是否大于堆顶，是的话，把堆顶pop出来，然后把新数据push下去
		if h.nums[0] >= val {
			return
		}
		h.Pop()
	}

	// 自下往上, 注意，前面Pop已经把h.size减1了
	h.nums[h.size] = val
	i := h.size

	// 子节点index=i, 父节点index为 (i-1)/2
	for (i-1)/2 >= 0 && h.nums[i] < h.nums[(i-1)/2] {
		h.nums[i], h.nums[(i-1)/2] = h.nums[(i-1)/2], h.nums[i]
		i = (i - 1) / 2
	}

	h.size++
}

func (h *MinKthHeap) Pop() int {
	if h.size == 0 {
		return -1
	}

	result := h.nums[0]
	if h.size > 1 {
		h.nums[0] = h.nums[h.size-1]
	}
	h.size--
	minHeapify(h.nums, h.size-1, 0)
	return result
}

/**
对index = i进行堆化, n是最后一个节点的坐标

建最小堆，从上往下，不断把大的数值调整下去
*/
func minHeapify(nums []int, n, i int) {

	if n <= i {
		return
	}

	for {
		maxPos := i
		// 父节点index为i, 子节点为2*i+1, 2*i+2
		if 2*i+1 <= n && nums[2*i+1] < nums[i] {
			maxPos = 2*i + 1
		}

		// 注意这里是把子节点较小的一个换上去
		if 2*i+2 <= n && nums[2*i+2] < nums[maxPos] {
			maxPos = 2*i + 2
		}

		if maxPos == i {
			break
		}

		nums[i], nums[maxPos] = nums[maxPos], nums[i]
		i = maxPos
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthLargestElementInAnArray(t *testing.T) {
	//findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)

	findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
}
