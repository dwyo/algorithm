package leetcode
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
// 
//
// 示例 1： 
//
// 输入：head = [1,3,2]
//输出：[2,3,1] 
//
// 
//
// 限制： 
//
// 0 <= 链表长度 <= 10000 
// Related Topics 链表 
// 👍 82 👎 0

type ListNode struct{
	Val int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	node := head
	cnt := 0
	for node != nil {
		cnt++
		node = node.Next
	}
	node = head
	ret := make([]int, cnt)
	for i := cnt-1; i>=0;i-- {
		ret[i] = node.Val
		node = node.Next
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
