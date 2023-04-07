package main

import (
	"testing"
)

/**
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -100 <= Node.val <= 100


 Related Topics 树 广度优先搜索 二叉树 👍 755 👎 0

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

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	var ans [][]int
	for len(stack) > 0 {
		curSize := len(stack)
		curRoundAns := make([]int, curSize)
		for i := 0; i < curSize; i++ {
			curRoundAns[i] = stack[i].Val

			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}

			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}

		stack = stack[curSize:]
		ans = append(ans, curRoundAns)
	}

	// 处理结果，奇数层调转
	for j := 0; j < len(ans); j++ {
		if j%2 == 0 {
			continue
		}

		curRoundAns := ans[j]
		for k := 0; k < len(curRoundAns)/2; k++ {
			curRoundAns[k], curRoundAns[len(curRoundAns)-k-1] = curRoundAns[len(curRoundAns)-k-1], curRoundAns[k]
		}
	}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeZigzagLevelOrderTraversal(t *testing.T) {

}
