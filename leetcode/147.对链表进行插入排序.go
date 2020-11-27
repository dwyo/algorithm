package leetcode

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	//
	lastSorted := head
	curr := head.Next
	for curr != nil {
		// 小于，则不用做操作，符合正序，进行下一位对比
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else { // 小于则交换位置
			// 和左侧的所有元素进行对比
			prev := dummyHead
			for prev.Next.Val < curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}

// @lc code=end
