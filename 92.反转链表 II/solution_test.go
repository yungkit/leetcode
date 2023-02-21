package main

import (
	"testing"
)

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节
点，返回 反转后的链表 。



 示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]


 示例 2：


输入：head = [5], left = 1, right = 1
输出：[5]




 提示：


 链表中节点数目为 n
 1 <= n <= 500
 -500 <= Node.val <= 500
 1 <= left <= right <= n




 进阶： 你可以使用一趟扫描完成反转吗？

 Related Topics 链表 👍 1500 👎 0

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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if left > right {
		return head
	}

	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	var firstPartEnd = dummyNode

	// 先找到left点切割点附近的两个点的位置
	for i := 0; i < left-1; i++ {
		firstPartEnd = firstPartEnd.Next
	}

	secondPartEnd := firstPartEnd.Next
	for i := 0; i < right-left; i++ {
		secondPartEnd = secondPartEnd.Next
	}

	secondPartStart := firstPartEnd.Next
	thirdPartStart := secondPartEnd.Next

	// 这里要注意先切断!!!
	firstPartEnd.Next = nil
	secondPartEnd.Next = nil

	// 把第二部分翻转
	reverseList(secondPartStart)

	// 把第一部分拼上原来的secondPartEnd， 原来的secondPartStart拼上thirdPartStart即可
	firstPartEnd.Next = secondPartEnd
	secondPartStart.Next = thirdPartStart

	return dummyNode.Next
}

func reverseList(head *ListNode) *ListNode {

	var prev, cur *ListNode = nil, head

	for cur != nil {
		tempNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = tempNext
	}

	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedListIi(t *testing.T) {

}
