package main

import (
	"testing"
)

/**
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

 你应当 保留 两个分区中每个节点的初始相对位置。



 示例 1：


输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]


 示例 2：


输入：head = [2,1], x = 2
输出：[1,2]




 提示：


 链表中节点的数目在范围 [0, 200] 内
 -100 <= Node.val <= 100
 -200 <= x <= 200


 Related Topics 链表 双指针 👍 683 👎 0

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

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var l1, l2 *ListNode = &ListNode{Val: 0}, &ListNode{Val: 0}

	cur := head
	p1, p2 := l1, l2
	for cur != nil {
		if cur.Val < x {
			p1.Next = cur
			//cur = cur.Next
			p1 = cur
		} else {
			p2.Next = cur
			p2 = cur
		}
		cur = cur.Next
	}

	p2.Next = nil
	p1.Next = l2.Next

	return l1.Next

}

//leetcode submit region end(Prohibit modification and deletion)

func TestPartitionList(t *testing.T) {

}
