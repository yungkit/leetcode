package main

import (
	"testing"
)

/**
给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。

 如果树中有不止一个众数，可以按 任意顺序 返回。

 假定 BST 满足如下定义：


 结点左子树中所含节点的值 小于等于 当前节点的值
 结点右子树中所含节点的值 大于等于 当前节点的值
 左子树和右子树都是二叉搜索树




 示例 1：


输入：root = [1,null,2,2]
输出：[2]


 示例 2：


输入：root = [0]
输出：[0]




 提示：


 树中节点的数目在范围 [1, 10⁴] 内
 -10⁵ <= Node.val <= 10⁵




 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 617 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var cur, count, maxCount int
	var ans []int
	update := func(x int) {
		if cur == x {
			count++
		} else {
			count = 1
			cur = x
		}

		if count == maxCount {
			ans = append(ans, x)
		} else if count > maxCount {
			maxCount = count
			ans = []int{x}
		}
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func findMode2(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	countMap := make(map[int]int)
	stack := []*TreeNode{root}

	for len(stack) > 0 {

		curRoundSize := len(stack)

		for i := 0; i < curRoundSize; i++ {
			cur := stack[i]
			val := cur.Val
			if _, ok := countMap[val]; !ok {
				countMap[val] = 0
			}

			countMap[val]++

			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}

			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}

		stack = stack[curRoundSize:]
	}

	maxValCount := 0
	for _, v := range countMap {
		if v > maxValCount {
			maxValCount = v
		}
	}

	ans := make([]int, 0)
	for k, v := range countMap {
		if v == maxValCount {
			ans = append(ans, k)
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindModeInBinarySearchTree(t *testing.T) {

}
