package main

import (
	"testing"
)

/**
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

 叶子节点 是指没有子节点的节点。



 示例 1：




输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]


 示例 2：




输入：root = [1,2,3], targetSum = 5
输出：[]


 示例 3：


输入：root = [1,2], targetSum = 0
输出：[]




 提示：


 树中节点总数在范围 [0, 5000] 内
 -1000 <= Node.val <= 1000
 -1000 <= targetSum <= 1000


 注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/

 Related Topics 树 深度优先搜索 回溯 二叉树 👍 402 👎 0

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

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	path := []int{}

	var dfs func(node *TreeNode, leftSum int)
	dfs = func(node *TreeNode, leftSum int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)

		// 应该是一直往叶子节点找, 符合条件就加入ans
		if node.Left == nil && node.Right == nil && node.Val == leftSum {
			ans = append(ans, append([]int{}, path...))
			// 访问到叶子节点了，要弹出自己
			path = path[:len(path)-1]
			return
		}

		dfs(node.Left, leftSum-node.Val)
		dfs(node.Right, leftSum-node.Val)

		// 说明这里也访问到叶子节点了，要弹出自己
		path = path[:len(path)-1]
	}

	dfs(root, target)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaShuZhongHeWeiMouYiZhiDeLuJingLcof(t *testing.T) {

}
