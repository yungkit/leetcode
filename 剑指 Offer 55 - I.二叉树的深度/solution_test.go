package main

import (
	"testing"
)

/**
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

 例如：

 给定二叉树 [3,9,20,null,null,15,7]，

     3
   / \
  9  20
    /  \
   15   7

 返回它的最大深度 3 。



 提示：


 节点总数 <= 10000


 注意：本题与主站 104 题相同：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 227 👎 0

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftChildDepth := maxDepth(root.Left) + 1
	rightChildDepth := maxDepth(root.Right) + 1

	if leftChildDepth > rightChildDepth {
		return leftChildDepth
	}

	return rightChildDepth

}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaShuDeShenDuLcof(t *testing.T) {

}
