package leetcode

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (62.13%)
 * Likes:    346
 * Dislikes: 0
 * Total Accepted:    81.7K
 * Total Submissions: 131.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
 *
 * 你应当保留两个分区中每个节点的初始相对位置。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：head = 1->4->3->2->5->2, x = 3
 * 输出：1->2->2->4->3->5
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	left := &ListNode{}
	lhead := left
	right := &ListNode{}
	rhead := right
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	right.Next = nil
	left.Next = rhead.Next
	return lhead.Next
}

// @lc code=end
