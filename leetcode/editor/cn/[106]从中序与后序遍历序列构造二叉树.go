package leetcode

//根据一棵树的中序遍历与后序遍历构造二叉树。
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
// Related Topics 树 深度优先搜索 数组 
// 👍 435 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	n := len(postorder)
	root := &TreeNode{postorder[n-1], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[n-1] {
			break
		}
	}
	root.Left = buildTree(postorder[:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(postorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
