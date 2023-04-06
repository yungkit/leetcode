package main

import (
	"testing"
)

/**
给你一个链表数组，每个链表都已经按升序排列。

 请你将所有链表合并到一个升序链表中，返回合并后的链表。



 示例 1：

 输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6


 示例 2：

 输入：lists = []
输出：[]


 示例 3：

 输入：lists = [[]]
输出：[]




 提示：


 k == lists.length
 0 <= k <= 10^4
 0 <= lists[i].length <= 500
 -10^4 <= lists[i][j] <= 10^4
 lists[i] 按 升序 排列
 lists[i].length 的总和不超过 10^4


 Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2397 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 递归思路1
//func mergeKLists(lists []*ListNode) *ListNode {
//
//	if len(lists) == 0 {
//		return nil
//	}
//
//	if len(lists) == 1 {
//		return lists[0]
//	}
//
//	if len(lists) == 2 {
//		return mergeTwoLists(lists[0], lists[1])
//	}
//
//	size := len(lists)
//	mid := size / 2
//	firstPartLists := make([]*ListNode, mid)
//	for i := 0; i < mid; i++ {
//		firstPartLists[i] = lists[i]
//	}
//
//	left := mergeKLists(firstPartLists)
//
//	secondPartLists := make([]*ListNode, size-mid)
//	for j := mid; j < size; j++ {
//		secondPartLists[j-mid] = lists[j]
//	}
//	right := mergeKLists(secondPartLists)
//
//	return mergeTwoLists(left, right)
//}
//
//func mergeTwoLists(l1, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//
//	if l2 == nil {
//		return l1
//	}
//
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	}
//
//	l2.Next = mergeTwoLists(l1, l2.Next)
//	return l2
//}

//递归思路2
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))

}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyNode := &ListNode{}
	cur := dummyNode
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}

		cur = cur.Next
	}

	if p1 != nil {
		cur.Next = p1
	}

	if p2 != nil {
		cur.Next = p2
	}

	return dummyNode.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeKSortedLists(t *testing.T) {

}
