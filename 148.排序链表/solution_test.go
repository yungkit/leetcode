package main

import (
	"testing"
)

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。






 示例 1：


输入：head = [4,2,1,3]
输出：[1,2,3,4]


 示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]


 示例 3：


输入：head = []
输出：[]




 提示：


 链表中节点的数目在范围 [0, 5 * 10⁴] 内
 -10⁵ <= Node.val <= 10⁵




 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

 Related Topics 链表 双指针 分治 排序 归并排序 👍 1917 👎 0

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

func sortList(head *ListNode) *ListNode {
	// 使用自顶向下的归并排序，这里先不管什么空间O(1), 面试的时候先把题目做出来
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head

	// 这里要注意结束条件不是nil, 而是当前递归的tail
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	return mergeList(sort(head, mid), sort(mid, tail))

}

/**
使用归并排序的方式合并两个链表
*/
func mergeList(l1, l2 *ListNode) *ListNode {

	dummyNode := &ListNode{Val: -1}

	temp, temp1, temp2 := dummyNode, l1, l2

	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}

		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	} else {
		temp.Next = temp2
	}

	return dummyNode.Next

}

//leetcode submit region end(Prohibit modification and deletion)

func TestSortList(t *testing.T) {

}
