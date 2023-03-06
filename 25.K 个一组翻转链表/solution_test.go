package main

import (
	"testing"
)

/**
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

 k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。



 示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]


 示例 2：




输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]



提示：


 链表中的节点数目为 n
 1 <= k <= n <= 5000
 0 <= Node.val <= 1000




 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？




 Related Topics 递归 链表 👍 1920 👎 0

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	end := dummyNode

	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		if end == nil {
			break
		}

		// 标记待翻转区域的第一个节点，跟未处理区域的第一个节点
		start := pre.Next
		next := end.Next

		// 断开待翻转区域
		end.Next = nil

		// 开始翻转，翻转后由pre连接
		pre.Next = reverseList(start)

		// 翻转好后的尾节点接上未处理区域
		start.Next = next

		// pre, end都切到未处理区域的第一个节点，等待下一轮的处理
		pre = start
		end = pre
	}

	return dummyNode.Next
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

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseNodesInKGroup(t *testing.T) {

}
