package leetcode

//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明: 
//1 ≤ m ≤ n ≤ 链表长度。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL 
// Related Topics 链表 
// 👍 632 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var temp *ListNode

	var reverseN func(*ListNode, int) *ListNode
	reverseN = func(node *ListNode, x int) *ListNode {
		if x == 1 {
			temp = node.Next
			return node
		}
		last := reverseN(node.Next, x-1)
		node.Next.Next = node
		node.Next = temp
		return last
	}
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
