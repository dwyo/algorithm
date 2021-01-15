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

func (t *BTree) Preorder(root *BTree)  {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	t.Preorder(root.Left)
	t.Preorder(root.Right)

}