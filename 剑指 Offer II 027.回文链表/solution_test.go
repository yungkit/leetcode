package main

import (
	"testing"
)

/**
给定一个链表的 头节点 head ，请判断其是否为回文链表。

 如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。



 示例 1：




输入: head = [1,2,3,3,2,1]
输出: true

 示例 2：




输入: head = [1,2]
输出: false




 提示：


 链表 L 的长度范围为 [1, 10⁵]
 0 <= node.val <= 9




 进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？




 注意：本题与主站 234 题相同：https://leetcode-cn.com/problems/palindrome-linked-list/

 Related Topics 栈 递归 链表 双指针 👍 93 👎 0

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

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到链表中间节点
	firstHalfEnd := endOfFirstHalf(head)

	// 翻转后半部分
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 双指针的方式比对对应顺序节点的值
	p1 := head
	p2 := secondHalfStart

	result := true
	// 后半部分会断一点，所以判断条件是p2
	for p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表后半部分
	firstHalfEnd.Next = reverseList(secondHalfStart)

	return result
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTemp
	}
	return pre
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAMhZSa(t *testing.T) {
	node4 := &ListNode{Val: 1, Next: nil}
	node3 := &ListNode{Val: 2, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	println(isPalindrome(node1))
}
