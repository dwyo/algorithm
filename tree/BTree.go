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
