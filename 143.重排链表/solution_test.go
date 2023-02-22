package main

import (
	"testing"
)

/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：


L0 → L1 → … → Ln - 1 → Ln


 请将其重新排列后变为：


L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



 示例 1：




输入：head = [1,2,3,4]
输出：[1,4,2,3]

 示例 2：




输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]



 提示：


 链表的长度范围为 [1, 5 * 10⁴]
 1 <= node.val <= 1000


 Related Topics 栈 递归 链表 双指针 👍 1140 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 找到中间点
	middleNode := findMiddleNode(head)

	secondHalfList := middleNode.Next

	// 切割链表，反转后半部分
	middleNode.Next = nil

	secondHalfList = reverseList(secondHalfList)

	// 合并两个链表
	mergeList(head, secondHalfList)

}

func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTemp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {

	var l1Temp, l2Temp *ListNode

	// 在纸上画画就能理解了，非常简单
	for l1 != nil && l2 != nil {
		l1Temp = l1.Next
		l2Temp = l2.Next

		l1.Next = l2
		l2.Next = l1Temp

		l1 = l1Temp
		l2 = l2Temp
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReorderList(t *testing.T) {

}
