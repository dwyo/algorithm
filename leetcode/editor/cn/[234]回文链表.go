package leetcode

//请判断一个链表是否为回文链表。
//
// 示例 1: 
//
// 输入: 1->2
//输出: false 
//
// 示例 2: 
//
// 输入: 1->2->2->1
//输出: true
// 
//
// 进阶： 
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
// Related Topics 链表 双指针 
// 👍 816 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	left := head
	slow, fast := head, head
	var reverse func(*ListNode) *ListNode
	reverse = func(node *ListNode) *ListNode {
		curr := node
		var prev *ListNode
		for curr != nil {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}
		return prev
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	right := reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

// 递归
//func isPalindrome(head *ListNode) bool {
//	left := head
//	var traverse func(*ListNode)bool
//	traverse = func(node *ListNode) bool {
//		if node == nil {
//			return true
//		}
//		// 本质上是模拟了一个栈，在拿出来和原链表比较
//		res := traverse(node.Next)
//		res = res && (node.Val == left.Val)
//		left = left.Next
//		return res
//	}
//	return traverse(head)
//}
