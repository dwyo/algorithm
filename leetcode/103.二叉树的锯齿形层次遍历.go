package leetcode

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (55.19%)
 * Likes:    322
 * Dislikes: 0
 * Total Accepted:    88K
 * Total Submissions: 158.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层序遍历如下：
 *
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 分析，把每一层的节点都放出来
	var res [][]int
	if root == nil {
		return res
	}
	quene := []*TreeNode{root}
	isLeftStart := true
	for len(quene) > 0 {
		l := quene
		n := len(l)
		s := make([]int, n)
		for i := 0; i < n; i++ {
			node := quene[i]
			if node.Left != nil {
				quene = append(quene, node.Left)
			}
			if node.Right != nil {
				quene = append(quene, node.Right)
			}
			if isLeftStart { // 左侧不做处理
				s[i] = node.Val
			} else { // 右侧，倒着插入
				s[n-i-1] = node.Val
			}
		}
		res = append(res, s)
		isLeftStart = !isLeftStart
		quene = quene[n:]
	}
	return res
}

// @lc code=end
