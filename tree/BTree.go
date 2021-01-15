package tree

import "fmt"

type BTree struct {
	Val   int
	Left  *BTree
	Right *BTree
}

func (t *BTree) add(val int) {

}

func (t *BTree) Find(needle int) *BTree {
	if t == nil {
		return nil
	}

	if needle == t.Val {
		return t
	}
	left := t.Left.Find(needle)
	if left.Val == needle {
		return left
	}
	right := t.Right.Find(needle)
	if right.Val == needle {
		return right
	}
	return nil
}

func (t *BTree) Remove(needle int) bool {

	return true
}

/**
前序遍历
*/
func (t *BTree) Preorder(root *BTree) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	t.Preorder(root.Left)
	t.Preorder(root.Right)
}

func (t *BTree) Inorder(root *BTree) {
	if root == nil {
		return
	}
	t.Inorder(root.Left)
	fmt.Println(root.Val)
	t.Inorder(root.Right)
}

func (t *BTree) Postorder(root *BTree) {
	if root == nil {
		return
	}
	t.Postorder(root.Left)
	fmt.Println(root.Val)
	t.Postorder(root.Right)
}

/**
层序遍历
*/
func (t *BTree) LevelOrder(root *BTree) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	queue := []*BTree{root}
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			n--
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			ret = append(ret, node.Val)
			queue = queue[1:]
		}
	}
	return ret
}

func (t *BTree) BuildTree(preorder []int, inorder []int) *BTree {
	if len(preorder) == 0 {
		return nil
	}
	root := &BTree{preorder[0], nil, nil}
	i := 0
	// 找到中序遍历的根节点位置
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	p := preorder[1:len(inorder[:i])+1]
	root.Left = t.BuildTree(p, inorder[:i])
	root.Right = t.BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
